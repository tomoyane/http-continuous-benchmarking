#!/bin/bash

set -eu

go test -race -covermode atomic -coverprofile=cover.out ./... && go tool cover -html=cover.out -o index.html

git config --global user.name tomoyane
git config --global user.email "${ADMIN_EMAIL}"
git remote set-url --push origin https://tomoyane:"${GITHUB_TOKEN}"@github.com/tomoyane/http-continuous-benchmarking.git
git remote -v

git fetch
git checkout coverage

is_changed=$(git status | grep "nothing to commit" | wc -l)
if [ "$is_changed" -eq 1 ]; then
  echo "Not need to update coverage"
else
  git add --force index.html
  git commit -m "[ci skip] Update coverage"
  git remote set-url --push origin https://tomoyane:${GITHUB_TOKEN}@github.com/tomoyane/http-continuous-benchmarking.git
  git push origin HEAD:coverage --force
fi
