FROM golang:1.24.1

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/software_api
COPY . $GOPATH/src/software_api
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./software_api"]
