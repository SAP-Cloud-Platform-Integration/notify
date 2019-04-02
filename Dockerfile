FROM golang:1.12-alpine

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

# start
CMD ["app", "start"]