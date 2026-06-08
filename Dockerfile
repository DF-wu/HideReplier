FROM node:20-alpine AS frontend-builder
WORKDIR /web

COPY web/package.json web/package-lock.json ./
RUN npm ci

COPY web/ ./
ARG VITE_BOT_VERSION="dev"
ENV VITE_BOT_VERSION=${VITE_BOT_VERSION}
RUN npm run build

FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/hidereplier-go ./cmd/server

FROM alpine:3.20
WORKDIR /app

COPY --from=builder /out/hidereplier-go /app/hidereplier-go
COPY --from=frontend-builder /web/dist /app/static
COPY --from=builder /app/src/main/resources/static/thumbs /app/static/thumbs

ENV PORT=8082
ENV STATIC_DIR=/app/static
EXPOSE 8082

ENTRYPOINT ["/app/hidereplier-go"]
