FROM golang:1.23.2-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -ldflags="-s -w" -o ./go-hoyo-daily ./cmd/bot/

FROM alpine:latest AS go-hoyo-daily
WORKDIR /app
COPY --from=builder /app/go-hoyo-daily .
ENTRYPOINT ["./go-hoyo-daily"]
