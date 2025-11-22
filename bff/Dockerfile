FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY bff/go.mod bff/go.sum ./
# We need to handle the local replace directive. 
# In a real scenario, we should copy the service modules too or use a workspace.
# Since we are in a monorepo, we can copy everything needed.

COPY user-service/ ./user-service/
COPY post-service/ ./post-service/
COPY bff/ ./bff/

WORKDIR /app/bff
RUN go mod download
RUN go build -o server server.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/bff/server .

EXPOSE 8080

CMD ["./server"]
