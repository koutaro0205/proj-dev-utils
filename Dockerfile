FROM golang:1.22.2-alpine

WORKDIR /app

COPY app/go.mod .
COPY app/go.sum .

RUN go mod download
