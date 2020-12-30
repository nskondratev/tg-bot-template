BINPATH   = $(PWD)/bin
LOCAL     = $(MODULE)
MODULE    = $(shell $(GO) list -m)
PATHS     = $(shell $(GO) list ./... 2> /dev/null | sed -e "s|$(MODULE)/\{0,1\}||g")
SHELL     = /bin/bash -euo pipefail

export PATH := $(BINPATH):$(PATH)

.PHONY: tools
tools:
	cd tools && go mod tidy && go mod verify && go generate tools.go

.PHONY: go-deps
go-deps:
	go mod tidy && go mod vendor && go mod verify

.PHONY: deps
deps: tools go-deps

.PHONY: generate
generate: go-generate

.PHONY: go-gen
go-gen:
	go generate ./...

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test:
	go test -race ./... -coverprofile=coverage.out && go tool cover -func=coverage.out && echo "Ok"

.PHONY: deploy
deploy:
	gcloud functions deploy BotUpdate --runtime go113 --trigger-http --timeout 300s --allow-unauthenticated --region europe-west3

.PHONY: run
run:
	go run cmd/bot/main.go
