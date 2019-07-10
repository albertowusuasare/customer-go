# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang v1.11 base image
FROM golang:1.12.6 as builder

# Add Maintainer Info
LABEL maintainer="Albert Owusu-Asare <albertowusuasa@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/albertowusuasare/customer-app

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
# https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
RUN go get -d -v ./...

# Install the package
RUN  CGO_ENABLED=0 GOOS=linux go install -v ./...


### seond stage
# deployment image
FROM scratch

# copy ca-certificates from builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

WORKDIR /bin/

COPY --from=builder /go/bin/customer-svc .

# This container exposes port 8080 to the outside world
EXPOSE 5090

# Run the executable
CMD ["./customer-svc"]