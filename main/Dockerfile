FROM golang:latest

RUN mkdir /go/src/app

WORKDIR /go/src/app

COPY . .

EXPOSE 8080

CMD ["go", "run", "main.go"]