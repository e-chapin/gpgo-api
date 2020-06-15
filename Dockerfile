#FROM heroku/heroku:18-build as build
#
#COPY . /app
#WORKDIR /app
#
## Setup buildpack
#RUN mkdir -p /tmp/buildpack/heroku/go /tmp/build_cache /tmp/env
#RUN curl https://codon-buildpacks.s3.amazonaws.com/buildpacks/heroku/go.tgz | tar xz -C /tmp/buildpack/heroku/go
#
##Execute Buildpack
#RUN STACK=heroku-18 /tmp/buildpack/heroku/go/bin/compile /app /tmp/build_cache /tmp/env
#
## Prepare final, minimal image
#FROM heroku/heroku:18
#
#COPY --from=build /app /app
#ENV HOME /app
#WORKDIR /app
#RUN useradd -m heroku
#USER heroku
#CMD /app/bin/gpgo


# Build Go module
FROM golang:1.14 AS gobuilder
WORKDIR /app
COPY ./config ./config
COPY ./controllers ./controllers
COPY ./db ./db
COPY ./helpers ./helpers
COPY ./models ./models
COPY ./server ./server
COPY ./static ./static
COPY ./views ./views
COPY ./cmd ./cmd
COPY ./cmd ./cmd
COPY config.go .
COPY go.mod .
COPY go.sum .
COPY main.go .
RUN go mod download
RUN go build -o bin/gpgo ./main.go

# Build Web resources
FROM node:13 AS webbuilder
WORKDIR /views
COPY /views .
#ENV REACT_APP_WEBSOCKET_ENDPOINT="wss://googlekeep-anselm94.herokuapp.com/query"
#RUN npm ci --only=production
RUN webpack build

# Build final image
# Need to use Golang image, as SQLite requires CGO,
# and cannot be created a standalone executable
FROM golang:1.13
WORKDIR /
COPY --from=gobuilder /app/bin/ ./
COPY --from=webbuilder /views/build ./static
COPY run.sh .
RUN apt-get update && apt-get install -y uuid-runtime
ENV HOST=https://milo-gpgo.herokuapp.com
EXPOSE 80
CMD ["sh", "run.sh"]