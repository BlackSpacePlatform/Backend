FROM golang

RUN mkdir -p /go/src/github.com/LensPlatform/BlackSpace/podinfo

ADD . /go/src/github.com/LensPlatform/BlackSpace/podinfo
WORKDIR /go/src/github.com/LensPlatform/BlackSpace/podinfo

RUN go get -t -v ./...
RUN go get github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT watcher -run /go/src/github.com/LensPlatform/BlackSpace/podinfo/cmd/podinfo -watch github.com/LensPlatform/BlackSpace/podinfo
