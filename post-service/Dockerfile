FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY post-service/go.mod post-service/go.sum ./
COPY post-service/ .
RUN go mod download
RUN go build -o server main.go

FROM alpine:latest

RUN apk add --no-cache tzdata

WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 50052

CMD ["./server"]
