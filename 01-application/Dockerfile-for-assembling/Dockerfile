FROM golang:latest

ENV AUTHOR=Razin
WORKDIR /go/src/app
COPY . .

RUN go build -o app main.go

CMD ["/go/src/app/app"]
