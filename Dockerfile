FROM golang:1.10.3-alpine3.8

LABEL MAINTAINER="mathnitin@gmail.com"

# Add webserver binary
RUN mkdir -p $GOPATH/src/github.com/mathnitin/ChuckNorris
ADD . $GOPATH/src/github.com/mathnitin/ChuckNorris

ENV GOBIN=/go/bin

RUN go install $GOPATH/src/github.com/mathnitin/ChuckNorris/main.go

EXPOSE 5000

CMD ["/go/bin/main"]
