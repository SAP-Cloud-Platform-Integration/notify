# build image
FROM golang:1.12-alpine3.9 AS build-env

# install build tools
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# build
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

# default env
ENV ENV_CONFIG true
ENV SMTP_PORT 465
ENV CPI_NAME cpi
ENV CONTACT_NAME admin

# distribution image
FROM alpine:3.9

# add CAs
RUN apk --no-cache add ca-certificates

WORKDIR /go/bin
COPY --from=build-env /go/bin /go/bin

# start
CMD ["./app", "start"]