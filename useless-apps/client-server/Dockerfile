FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go mod init jokeserver && go build -o main .

EXPOSE 9080

CMD ["./main"]