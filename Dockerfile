# Compile stage
FROM golang:latest

ENV CGO_ENABLED 0

ENV GOBIN /go/bin

ADD . /go/src/users
WORKDIR /go/src/users

# Go dep!
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

# The -gcflags "all=-N -l" flag helps us get a better debug experience
RUN go build -gcflags "all=-N -l" -o /server users

# Compile Delve
#RUN apk add --no-cache git
RUN go get github.com/derekparker/delve/cmd/dlv

FROM alpine

# Port 3000 belongs to our application, 40000 belongs to Delve
EXPOSE 8080 40000

# Allow delve to run on Alpine based containers.
RUN apk add --no-cache libc6-compat
RUN apk --update add ca-certificates

WORKDIR /

COPY --from=build-env /server /
COPY --from=build-env /go/bin/dlv /

# Run delve
CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "exec", "/server"]
