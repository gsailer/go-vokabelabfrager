FROM golang:1.7

MAINTAINER sublinus "sublinus@riseup.net"

ADD . /go/src/github.com/sublinus/go-vokabelabfrager
RUN go install github.com/sublinus/go-vokabelabfrager
RUN go get -v -d github.com/sublinus/go-vokabelabfrager

RUN mkdir /vokabelabfrager

ENTRYPOINT /go/bin/go-vokabelabfrager
EXPOSE 8080
