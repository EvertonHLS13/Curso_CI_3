FROM golang:1.18 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .


FROM ubuntu:latest

WORKDIR /app

COPY --from=builder /app/main main

EXPOSE 8000

CMD ["./main"]
