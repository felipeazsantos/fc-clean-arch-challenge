FROM golang:1.24.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/app ./cmd/servers/main.go

RUN ls -la /app/

FROM alpine:latest

WORKDIR /app

RUN mkdir -p /app/cmd/servers

COPY --from=builder /app/app /app/app

RUN echo 'db_driver=mysql' > /app/cmd/servers/.env && \
    echo 'db_host=mysql' >> /app/cmd/servers/.env && \
    echo 'db_port=3306' >> /app/cmd/servers/.env && \
    echo 'db_user=root' >> /app/cmd/servers/.env && \
    echo 'db_pass=root' >> /app/cmd/servers/.env && \
    echo 'db_name=orders' >> /app/cmd/servers/.env && \
    echo 'web_server_port=:8000' >> /app/cmd/servers/.env && \
    echo 'grpc_server_port=50051' >> /app/cmd/servers/.env && \
    echo 'graphql_server_port=8080' >> /app/cmd/servers/.env

RUN chmod +x /app/app

CMD ["/app/app"]
