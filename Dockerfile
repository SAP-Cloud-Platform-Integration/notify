FROM golang:1.12-alpine

# install git
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]