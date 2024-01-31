FROM golang:1.21-alpine

WORKDIR $GOPATH/src

COPY . .

RUN go mod download
RUN go build -o /cvg-server

EXPOSE 8080

CMD ["/cvg-server"]