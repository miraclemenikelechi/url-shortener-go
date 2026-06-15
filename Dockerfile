# stage 1: Build the Go application
FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /url-shortner-go

# stage 2: Run the application
FROM alpine

COPY --from=builder /url-shortner-go /url-shortner-go

EXPOSE 8649

CMD ["/url-shortner-go"]
