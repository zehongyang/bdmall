FROM golang:1.12-alpine
ENV GO111MODULE on
ENV GOPROXY https://goproxy.io
COPY . /go/src/bdmall
RUN cd /go/src/bdmall && go build -o bdmall
