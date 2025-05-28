# build stage
FROM golang:alpine AS builder

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o car-service ./cmd/main.go

# final image
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/car-service .
CMD ["./car-service"]
