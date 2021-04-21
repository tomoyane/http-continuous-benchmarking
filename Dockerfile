FROM golang:1.15

RUN mkdir /tomoyane
WORKDIR /tomoyane
RUN go build
COPY ./http-continuous-benchmarking /tomoyane
RUN chmod +x /tomoyane/http-continuous-benchmarking

ENTRYPOINT ["./http-continuous-benchmarking"]
