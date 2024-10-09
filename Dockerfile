FROM golang:1.20-alpine

WORKDIR /app

COPY . .

WORKDIR /app/app

RUN go mod tidy
RUN go build -o distributed-file-storage .

CMD ["/app/app/distributed-file-storage"]
