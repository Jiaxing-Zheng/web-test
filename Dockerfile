FROM golang:latest

WORKDIR $GOPATH/src/test
COPY . $GOPATH/src/test
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./test"]