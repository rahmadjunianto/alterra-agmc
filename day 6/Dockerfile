# step 1: build executable binary
FROM golang:1.19-alpine AS builder
LABEL maintainer="Rahmad Junianto<rahmadjunianto.rj@gmail.com>"
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o /agmc

# step 2: build a small image
FROM alpine:3.16.0
WORKDIR /app
COPY --from=builder agmc .
COPY .env .
EXPOSE 3000
CMD ["./agmc","-m=migrate"]