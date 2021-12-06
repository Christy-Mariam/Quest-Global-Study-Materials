FROM golang:latest

COPY test.go .

RUN go build test.go

EXPOSE 32100

CMD ["./test"]

