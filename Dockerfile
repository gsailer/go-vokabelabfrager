FROM scratch

MAINTAINER sublinus

ADD https://github.com/sublinus/go-vokabelabfrager.git main /

ENTRYPOINT ["/main"]

EXPOSE 8080