FROM golang:1.15

RUN mkdir /tomoyane
WORKDIR /tomoyane
COPY . /tomoyane
RUN go build
ENTRYPOINT ["./http-continuous-benchmarking"]
