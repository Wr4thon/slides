# syntax = docker/dockerfile:1.0-experimental

FROM golang:alpine AS present

ARG PROJECT_PATH=/go/src/github.com/wr4thon/slides
ENV PROJECT_PATH=$PROJECT_PATH

RUN mkdir -p ${PROJECT_PATH}

WORKDIR ${PROJECT_PATH}

RUN apk --no-cache --update add tzdata build-base git ca-certificates

ENV TZ Europe/Berlin

FROM present 

EXPOSE 3999/tcp
COPY . .

RUN go get -u golang.org/x/tools/cmd/present  
RUN go install golang.org/x/tools/cmd/present@latest

ENTRYPOINT [ "present", "--http=0.0.0.0:3999", "-play", "-use_playground", "${PROJECT_PATH}" ]
