FROM golang:1.15

RUN mkdir /tomoyane
WORKDIR /tomoyane
COPY ./http-continuous-benchmarking /tomoyane
RUN chmod 755 /tomoyane/http-continuous-benchmarking

ENTRYPOINT ["./http-continuous-benchmarking"]
