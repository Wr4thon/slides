# syntax = docker/dockerfile:1.0-experimental

FROM golang:alpine AS present

ARG PROJECT_PATH=/slides
ENV PROJECT_PATH=$PROJECT_PATH

RUN apk --no-cache --update add tzdata build-base git ca-certificates

RUN go install golang.org/x/tools/cmd/present@latest

RUN mkdir -p ${PROJECT_PATH}

WORKDIR ${PROJECT_PATH}

ENV TZ Europe/Berlin

EXPOSE 80/tcp

ENTRYPOINT [ "present", "--http=0.0.0.0:80", "-play", "-use_playground", "-notes", "${PROJECT_PATH}" ]


FROM present

COPY . .
