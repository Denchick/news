FROM golang:latest

COPY . /go/src/app

WORKDIR /go/src/app/cmd/rssparser

RUN go build -o rssparser main.go

CMD ["./rssparser"]