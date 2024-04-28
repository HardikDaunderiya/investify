FROM golang:1.22-alpine3.19 AS builder
WORKDIR /

COPY . .
RUN go build -o main main.go

FROM alpine:3.19
