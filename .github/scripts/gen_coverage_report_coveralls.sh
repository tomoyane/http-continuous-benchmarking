#!/bin/bash

set -e -u -x

go test -race -covermode atomic -coverprofile=profile.cov ./...
GO111MODULE=off go get github.com/mattn/goveralls
$(go env GOPATH)/bin/goveralls -coverprofile=profile.cov -service=github
