FROM golang:1.13.5

ADD . /go/src/github.com/dmitry-msk777/gRPC_example

RUN go get "google.golang.org/grpc"
RUN go get "google.golang.org/grpc/grpclog"

RUN go install github.com/dmitry-msk777/gRPC_example

ENTRYPOINT ["/go/bin/dmitry777"]

EXPOSE 18181:18184