.PHONY: default help build package push helm run test clean

SHELL         = /bin/bash
APP_NAME      = teler
VERSION       = $(shell git describe --always --tags)
SQUAD         = infosec

default: help

help:
	@echo 'Management commands for ${APP_NAME}:'
	@echo
	@echo 'Usage:'
	@echo '    make build            Compile the project.'
	@echo '    make test             Run tests on a compiled project.'
	@echo '    make lint             Run linters on the project.'
	@echo '    make clean            Clean the directory tree.'

	@echo

build:
	@echo "Building ${APP_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags "-s -w -X ktbs.dev/teler/common.Version=${VERSION}" -buildvcs=false -o ./bin/${APP_NAME} ./cmd/${APP_NAME}

test:
	@echo "Testing ${APP_NAME} ${VERSION}"
	go test ./...

lint: golangci-lint semgrep

clean:
	@echo "Removing ${APP_NAME} ${VERSION}"
	@test ! -e ${APP_NAME} || rm ${APP_NAME}

golangci-lint:
	@echo "Run GolangCI-Lint"
	@if ! [ -x "$(shell command -v golangci-lint)" ]; then \
		echo "Error: GolangCI-Lint is not installed." >&2; \
		echo "       Get at https://github.com/golangci/golangci-lint/releases" >&2; \
		exit 1; \
	fi; \
	golangci-lint run ./... --tests=0 --issues-exit-code=1 -v

semgrep:
	@echo "Run semgrep"
	@if ! [ -x "$(shell command -v semgrep)" ]; then \
		echo "Error: semgrep is not installed." >&2; \
		echo "       See https://semgrep.dev/docs/getting-started/#run-semgrep-locally" >&2; \
		exit 1; \
	fi; \
	semgrep scan --config auto --dryrun -q --include "**.go"