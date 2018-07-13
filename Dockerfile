FROM golang:latest

MAINTAINER L "curmido@gmail.com"

WORKDIR $GOPATH/src/go_wechat
ADD . $GOPATH/src/go_wechat
RUN go build .

EXPOSE 9000

ENTRYPOINT ["./go_wechat"]