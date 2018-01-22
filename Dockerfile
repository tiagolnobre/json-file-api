FROM golang:1.9.2-alpine3.7
MAINTAINER tiago <tiago.l.nobre@gmail.com>

RUN apk add --update --no-cache git && \
  rm -rf /tmp/* /var/cache/apk/*

ARG app_env
ENV APP_ENV $app_env

COPY . /go/src/github.com/tiagolnobre/json-file-api
WORKDIR /go/src/github.com/tiagolnobre/json-file-api

RUN go get ./
RUN go build

CMD if [ ${APP_ENV} = production ]; \
      then \
      json-file-api; \
      else \
      go get github.com/pilu/fresh && \
      fresh; \
      fi

EXPOSE 3000
