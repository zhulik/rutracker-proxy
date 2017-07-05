FROM golang:1.8.3

LABEL maintainer="zhulik.gleb@gmail.com"

ENV WORKDIR /go/src/github.com/zhulik/rutracker-proxy/
RUN mkdir -p $WORKDIR
WORKDIR $WORKDIR

ADD . ./
RUN go get && go build

EXPOSE 8080

CMD ./rutracker-proxy