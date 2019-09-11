FROM golang:latest

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go get github.com/vivevincere/authserver
RUN go build -o main .

CMD ["/app/main"]
