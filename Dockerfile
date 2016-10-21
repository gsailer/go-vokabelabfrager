FROM golang:1.7

MAINTAINER sublinus "sublinus@riseup.net"

ADD . /go/src/github.com/sublinus/go-vokabelabfrager

RUN go get -v -d github.com/sublinus/go-vokabelabfrager

RUN go install github.com/sublinus/go-vokabelabfrager

RUN mkdir /usr/local/etc/vokabelabfrager

# temporary solution to provide a test data set
ADD ./data.db /usr/local/etc/vokabelabfrager/data.db

# ENTRYPOINT /go/bin/go-vokabelabfrager
CMD ["/go/bin/go-vokabelabfrager"]

# built Dockerfile for Heroku
# EXPOSE 8080
