# Build the Go API
FROM golang:1.14
WORKDIR /app
RUN mkdir bin
ADD src src
WORKDIR /app/src

RUN ls -lah
WORKDIR /app/src/static
RUN ls -lah

WORKDIR /app/src


RUN go clean -modcache

ENV GO111MODULE=on

# sync vendoring
RUN go mod vendor
# download packages
RUN go mod download
RUN go mod verify
RUN go build -v -o ../bin/gpgo

RUN ls -lah
WORKDIR /app
RUN ls -lah
ENTRYPOINT ["./bin/gpgo"]