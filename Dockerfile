FROM golang:1.6

RUN mkdir -p /go/src/github.com/jessemillar
ADD . /go/src/github.com/jessemillar/stalks

WORKDIR /go/src/github.com/jessemillar/stalks
RUN go get -d -v
RUN go install -v

CMD ["/go/bin/stalks"]

EXPOSE 15000
