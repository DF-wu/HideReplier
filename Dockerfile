FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/hidereplier-go ./cmd/server

FROM alpine:3.20
WORKDIR /app

COPY --from=builder /out/hidereplier-go /app/hidereplier-go
COPY --from=builder /app/src/main/resources/static /app/src/main/resources/static

ENV PORT=8082
ENV STATIC_DIR=/app/src/main/resources/static
EXPOSE 8082

ENTRYPOINT ["/app/hidereplier-go"]
