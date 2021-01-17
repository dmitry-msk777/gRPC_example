FROM golang:1.13.5

ADD . /go/src/github.com/dmitry-msk777/gRPC_example

RUN go get "google.golang.org/grpc"
RUN go get "google.golang.org/grpc/grpclog"

RUN go install github.com/dmitry-msk777/gRPC_example

ENTRYPOINT /go/bin/gRPC_example

RUN chmod +x entrypoint.sh

EXPOSE 5300:5300

# https://temofeev.ru/info/articles/kak-sozdat-prostoy-mikroservis-na-golang-i-grpc-i-vypolnit-ego-konteynerizatsiyu-s-pomoshchyu-docker/
# https://habr.com/ru/post/238473/

# docker build -t grpc_example .
# docker run --publish 18184:5300 --name test --rm grpc_example

# https://hub.docker.com/repository/docker/dmitrymsk777/megaparsec_300/builds
# docker run dmitrymsk777/megaparsec_300
