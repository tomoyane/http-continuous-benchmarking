#!/bin/bash

set -eu

GOOS=linux GOARCH=amd64 go build
docker login -u "${DOCKER_USER}" -p "${DOCKER_PASSWORD}"

docker build -t tomohito/http-continuous-benckmarking:"${TAG_NAME}" .
docker push tomohito/http-continuous-benckmarking:"${TAG_NAME}"
