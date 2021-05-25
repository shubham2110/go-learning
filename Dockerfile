from golang:latest

RUN go get golang.org/x/tools/cmd/godoc
RUN go get -u github.com/kisielk/errcheck

EXPOSE 6060

CMD /go/bin/godoc -http=:6060


