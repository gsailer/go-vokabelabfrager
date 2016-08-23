FROM golang:1.7

MAINTAINER sublinus

ADD . /go/src/github.com/sublinus/go-vokabelabfrager
RUN go install github.com/sublinus/go-vokabelabfrager
RUN go get -v -d github.com/sublinus/go-vokabelabfrager

RUN mkdir /usr/local/etc/vokabelabfrager

ENTRYPOINT /go/bin/go-vokabelabfrager
EXPOSE 80
