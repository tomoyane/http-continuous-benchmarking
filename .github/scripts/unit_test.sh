#!/bin/bash

set -eu

go test -race -covermode atomic -coverprofile=profile.cov ./...
