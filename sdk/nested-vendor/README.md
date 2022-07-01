# Nested vendor of the same sdk but using interfaces

```sh
$ cd sdk/nested-vendor/entrypoint
$ DOCKER_BUILDKIT=0 docker build -t entrypoint-example .
Sending build context to Docker daemon  23.04kB
Step 1/6 : from golang
 ---> 46ae95f04a69
Step 2/6 : WORKDIR /go/src/entrypoint
 ---> Using cache
 ---> 89a648e8d6b3
Step 3/6 : COPY . ./
 ---> 2a9ec83c0471
Step 4/6 : COPY handler /go/src/
 ---> b6a0a7fef65b
Step 5/6 : RUN pwd && ls -al .
 ---> Running in 1ee8a5393ad6
/go/src/entrypoint
total 32
drwxr-xr-x 1 root root 4096 Jul  1 13:43 .
drwxrwxrwx 1 root root 4096 Jul  1 13:43 ..
-rw-rw-r-- 1 root root  127 Jul  1 13:40 Dockerfile
-rw-rw-r-- 1 root root  220 Jul  1 13:42 go.mod
-rw-rw-r-- 1 root root  313 Jul  1 13:43 go.sum
drwxrwxr-x 3 root root 4096 Jul  1 13:42 handler
-rw-rw-r-- 1 root root  366 Jul  1 13:40 main.go
drwxrwxr-x 4 root root 4096 Jul  1 13:43 vendor
Removing intermediate container 1ee8a5393ad6
 ---> 3887019b97b2
Step 6/6 : RUN GO111MODULE=off go run main.go
 ---> Running in 582cdbd81ce5
hi
Removing intermediate container 582cdbd81ce5
 ---> d59bfe2120cb
Successfully built d59bfe2120cb
Successfully tagged entrypoint-example:latest
```
