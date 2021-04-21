FROM golang:1.15

RUN mkdir /tomoyane
WORKDIR /tomoyane
COPY . /tomoyane
RUN GOOS=linux GOARCH=amd64 go build
ENTRYPOINT ["./http-continuous-benchmarking"]
