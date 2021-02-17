FROM golang:1.16

COPY . /go/src/app

WORKDIR /go/src/app/cmd/apiserver

RUN go build -o apiserver main.go

CMD ["./apiserver"]