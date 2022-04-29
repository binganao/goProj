FROM golang:1.18-alpine AS builder

WORKDIR /work

ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-s -w' -o go-shop main.go

# FROM debian:stable-slim
FROM alpine
WORKDIR /
COPY --from=builder /work/go-shop .
COPY config.yml ./

ENTRYPOINT [ "./go-shop" ]
