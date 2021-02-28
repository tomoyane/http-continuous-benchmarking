#!/bin/bash

set -eu

GOOS=linux GOARCH=amd64 go build
