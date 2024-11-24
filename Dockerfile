FROM golang:alpine

ENV GO111MODULE=auto \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

  

COPY . /grpcchat/


WORKDIR /grpcchat


RUN go mod download


EXPOSE 8080

CMD go run main.go