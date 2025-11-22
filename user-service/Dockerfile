FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY user-service/go.mod user-service/go.sum ./
COPY user-service/ .
RUN go mod download
RUN go build -o server main.go

FROM alpine:latest

RUN apk add --no-cache tzdata

WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 50051

CMD ["./server"]
