# stage 1: Build the Go application
FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /url-shortener-go

# stage 2: Run the application
FROM alpine

COPY --from=builder /app/memory/migrations /memory/migrations

COPY --from=builder /url-shortener-go /url-shortener-go

EXPOSE 8649

CMD ["/url-shortener-go"]
