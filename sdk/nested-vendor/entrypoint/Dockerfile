from golang

WORKDIR /go/src/entrypoint
COPY . ./
COPY handler /go/src/

RUN pwd && ls -al .
RUN GO111MODULE=off go run main.go