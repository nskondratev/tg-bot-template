// +build tools

package tools

import (
	_ "github.com/gojuno/minimock/v3/cmd/minimock"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
)

//go:generate go build -v -o=../bin/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint
//go:generate go build -v -o=../bin/minimock github.com/gojuno/minimock/v3/cmd/minimock
