FROM golang:latest

COPY test.go .

RUN go build test.go

EXPOSE 9055

CMD ["./test"]

