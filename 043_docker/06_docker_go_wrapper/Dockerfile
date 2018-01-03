# A hello world example with Go
FROM golang:1.9.2


RUN mkdir -p /go/srcgithub.com/w-k-s/golangwebdev/043_docker/06_docker_go_wrapper
WORKDIR /go/src/github.com/w-k-s/golangwebdev/043_docker/06_docker_go_wrapper

COPY . /go/src/github.com/w-k-s/golangwebdev/043_docker/06_docker_go_wrapper
RUN go-wrapper download && go-wrapper install

CMD ["go-wrapper", "run","-port=8080"]

EXPOSE 8080