FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o stress_test .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/ .

ENTRYPOINT ["/app/stress_test"]