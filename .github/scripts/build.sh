#!/bin/bash

set -eu

GOOS=linux GOARCH=amd64 go build
docker login -u "${DOCKER_USER}" -p "${DOCKER_PASSWORD}"

now=$(date "+%Y%m%d%H%M%S")date "+%Y%m%d%H%M%S"
docker build -t tomohito/http-continuous-benckmarking:"${now}" .
docker push tomohito/http-continuous-benckmarking:"${now}"

docker build -t tomohito/http-continuous-benckmarking:latest .
docker push tomohito/http-continuous-benckmarking:latest
