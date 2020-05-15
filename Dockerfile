# build image
FROM golang:1.12-alpine3.9 AS build-env

# install build tools
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# build
WORKDIR /app
COPY . .
RUN go build -mod=vendor .

# default env
ENV ENV_CONFIG true
ENV SMTP_PORT 465
ENV CPI_NAME cpi
ENV CONTACT_NAME admin

# distribution image
FROM alpine:3.9

# add CAs
RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=build-env /app/app /app/app

# start
CMD ["./app", "start"]