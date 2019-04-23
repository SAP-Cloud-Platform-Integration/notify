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

WORKDIR /go/src/app
COPY --from=build-env /go/src/app /go/src/app

# start
CMD ["app", "start"]