#!/bin/bash

if [ -z "$1" ]; then
  echo "You must provide new name for the module, e.g. ./scripts/rename.sh github.com/user/new-bot"
  exit 1
fi

NEW_NAME=$1

echo "Renaming module to $NEW_NAME..."

go mod edit -module $NEW_NAME go.mod
go mod edit -module $NEW_NAME/tools tools/go.mod

find . -type f \
    -name '*.go' \
    -exec sed -i -e "s,github.com/nskondratev/tg-bot-template,$NEW_NAME,g" {} \;

find . -type f -name '*.go-e' -delete

echo "Done"
