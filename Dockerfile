FROM golang:1.13-alpine

RUN apk update
RUN apk add fortune

WORKDIR /$GOPATH/src/github.com/flaviuvadan/fortune
COPY . ./

RUN go build

EXPOSE 8080

ENTRYPOINT ["go", "run", "."]