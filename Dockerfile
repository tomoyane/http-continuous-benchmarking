FROM golang:1.15

RUN mkdir /tomoyane
WORKDIR /tomoyane
COPY . /tomoyane
RUN cd /tomoyane && GOOS=linux GOARCH=amd64 go build
ENTRYPOINT ["./entrypoint.sh"]
