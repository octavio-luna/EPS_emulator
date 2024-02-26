FROM golang:1.16.3-alpine3.13 as builder
COPY . /app
WORKDIR /app
RUN go build -o main .
VOLUME /tmp
ENTRYPOINT [ "/app/main" ]