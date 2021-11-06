FROM golang:1.7.6-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o apiserver

EXPOSE 8000

CMD ["/apiserver"]