# Builder stage
FROM golang:1.24-alpine3.20 AS builder
WORKDIR /go/src/reverse-proxy
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go

# Final image
FROM alpine:latest
EXPOSE 80

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app
COPY --from=builder /go/src/reverse-proxy/main .

CMD ["./main"]
