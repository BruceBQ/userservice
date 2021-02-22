FROM golang:1.15-alpine

COPY . /userservice
WORKDIR /userservice

ENV GO111MODULE=on

RUN go build -o userservice "cmd/main.go"

CMD ["/userservice/userservice"]