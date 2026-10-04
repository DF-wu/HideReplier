# ---- frontend ---------------------------------------------------------------
FROM node:20-alpine AS frontend-builder
WORKDIR /web

COPY web/package.json web/package-lock.json ./
RUN npm ci --no-audit --no-fund

COPY web/ ./
ARG VITE_BOT_VERSION="dev"
ENV VITE_BOT_VERSION=${VITE_BOT_VERSION}
RUN npm run build

# ---- backend ----------------------------------------------------------------
FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Only the Go sources are needed here; keeping the context small gives better
# layer cache hits and avoids pulling node_modules Go packages into the build.
COPY cmd/ ./cmd/
COPY internal/ ./internal/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /out/hidereplier-go ./cmd/server

# ---- static asset packaging -------------------------------------------------
# Precompress text assets once at build time. The Go server sends the ".gz"
# sibling when the client accepts gzip, so no CPU is spent compressing at
# request time on the tiny production VM. Fly's edge proxy does not compress.
FROM alpine:3.20 AS assets
WORKDIR /static
COPY --from=frontend-builder /web/dist ./
COPY src/main/resources/static/thumbs ./thumbs
RUN find . -type f \( -name '*.js' -o -name '*.css' -o -name '*.svg' -o -name '*.html' -o -name '*.json' -o -name '*.txt' -o -name '*.map' \) \
      -exec sh -c 'gzip -9 -c "$1" > "$1.gz"' _ {} \;

# ---- runtime ----------------------------------------------------------------
FROM alpine:3.20
WORKDIR /app

RUN addgroup -S app && adduser -S -G app app

COPY --from=builder /out/hidereplier-go /app/hidereplier-go
COPY --from=assets /static /app/static

ENV PORT=8082
ENV STATIC_DIR=/app/static
# Make the Go GC work towards the 256MB VM limit instead of discovering it
# via the OOM killer. Fly's own [env] can override this.
ENV GOMEMLIMIT=192MiB
EXPOSE 8082

USER app
ENTRYPOINT ["/app/hidereplier-go"]
