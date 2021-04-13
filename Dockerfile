from golang:latest

RUN go get golang.org/x/tools/cmd/godoc

EXPOSE 6060

CMD /go/bin/godoc -http=:6060


