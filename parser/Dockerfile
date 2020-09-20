FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app
## Add this go mod download command to pull in any dependencies
RUN go mod download
## Our project will now successfully build with the necessary go libraries included.
RUN go build -o ./parser/main ./parser
## Our start command which kicks off
## our newly created binary executable
WORKDIR /app/parser
CMD ["/app/parser/main"]