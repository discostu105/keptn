# Use the offical Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.12 as builder
ARG version=develop

WORKDIR /go/src/github.com/keptn/keptn/mongodb-datastore

# Force the go compiler to use modules 
ENV GO111MODULE=on
ENV BUILDFLAGS=""

# Copy `go.mod` for definitions and `go.sum` to invalidate the next layer
# in case of a change in the dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

ARG debugBuild

# set buildflags for debug build
RUN if [ ! -z "$debugBuild" ]; then export BUILDFLAGS='-gcflags "all=-N -l"'; fi

# Copy local code to the container image. 
COPY . .

# Build the command inside the container.
# (You may fetch or manage dependencies here, either manually or with a tool like "godep".)
RUN cd cmd/mongodb-datastore-server && CGO_ENABLED=0 GOOS=linux go build $BUILDFLAGS -v -o mongodb-datastore

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:3.7
RUN apk add --no-cache ca-certificates

ARG debugBuild

# IF we are debugging, we need to install libc6-compat for delve to work on alpine based containers
RUN if [ ! -z "$debugBuild" ]; then apk add --no-cache libc6-compat; fi

# Copy the binary to the production image from the builder stage.
COPY --from=builder /go/src/github.com/keptn/keptn/mongodb-datastore/cmd/mongodb-datastore-server/mongodb-datastore /mongodb-datastore
COPY --from=builder /go/src/github.com/keptn/keptn/mongodb-datastore/swagger-ui /swagger-ui
COPY --from=builder /go/src/github.com/keptn/keptn/mongodb-datastore/swagger.yaml /swagger-ui/swagger.yaml

EXPOSE 8080

# required for external tools to detect this as a go binary
ENV GOTRACEBACK=all

# KEEP THE FOLLOWING LINES COMMENTED OUT!!! (they will be included within the travis-ci build)
#travis-uncomment ADD MANIFEST /
#travis-uncomment COPY entrypoint.sh /
#travis-uncomment ENTRYPOINT ["/entrypoint.sh"]

# Run the web service on container startup.
CMD ["/mongodb-datastore", "--port=8080", "--host=0.0.0.0"]
