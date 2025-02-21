FROM golang:latest AS builder

WORKDIR /app

RUN go mod init github.com/caleosun123/goproj
RUN go mod tidy

COPY . .

RUN go build -o main .

FROM debian:latest

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
