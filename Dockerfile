# Build stage 
FROM  golang:1.19-alpine3.17 AS builder

WORKDIR /app

COPY . .

RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

# Run stage 
FROM alpine:3.17
WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/migrate .

COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migrations ./migration

EXPOSE 8081

CMD ["/app/main"]

ENTRYPOINT [ "/app/start.sh" ]