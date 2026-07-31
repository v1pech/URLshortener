FROM golang:1.25.1-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o app-bin ./cmd

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app-bin .
COPY --from=builder /app/config/config.yaml ./config.yaml 
CMD ["./app-bin"]