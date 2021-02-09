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
	@echo '    make push             Push tagged images to registry.'
	@echo '    make run ARGS=        Run with supplied arguments.'
	@echo '    make test             Run tests on a compiled project.'
	@echo '    make test-extra       Run tests and run GolangCI-Lint.'
	@echo '    make clean            Clean the directory tree.'

	@echo

build:
	@echo "Building ${APP_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags "-s -w -X ktbs.dev/teler/common.Version=${VERSION}" -o ./bin/${APP_NAME} ./cmd/${APP_NAME}

run: build
	@echo "Running ${APP_NAME} ${VERSION}"
	${APP_NAME} ${ARGS}

test:
	@echo "Testing ${APP_NAME} ${VERSION}"
	go test ./...

test-extra: golangci-lint test

clean:
	@echo "Removing ${APP_NAME} ${VERSION}"
	@test ! -e ${APP_NAME} || rm ${APP_NAME}

test-with-report:
	@echo "Run Go Test with Report"
	go test ./... -json > test-reports.out -coverprofile=coverage-reports.out

golangci-lint:
	@echo "Run GolangCI-Lint"
	@if [ ! -d /tmp/golangci-lint ]; then \
		curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s v1.21.0; \
		mkdir -p /tmp/golangci-lint/; \
		mv ./bin/golangci-lint /tmp/golangci-lint/golangci-lint; \
	fi; \
	/tmp/golangci-lint/golangci-lint run ./... --issues-exit-code=1 -v

golangci-lint-with-report:
	@echo "Run GolangCI-Lint with Report"
	@if [ ! -d /tmp/golangci-lint ]; then \
		curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s v1.21.0; \
		mkdir -p /tmp/golangci-lint/; \
		mv ./bin/golangci-lint /tmp/golangci-lint/golangci-lint; \
	fi; \
	/tmp/golangci-lint/golangci-lint run ./... -v --out-format checkstyle > golangci-lint-reports.xml

sonarqube:
	@echo "Run SonarQube"
	@export SONAR_SCANNER_OPTS="-Xmx2048m"
	@if [ ! -d /tmp/sonar-scanner ]; then \
		wget https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-4.0.0.1744-linux.zip; \
		unzip sonar-scanner-cli-4.0.0.1744-linux.zip; \
		rm sonar-scanner-cli-4.0.0.1744-linux.zip; \
		mv sonar-scanner-4.0.0.1744-linux /tmp/sonar-scanner; \
	fi; \
	/tmp/sonar-scanner/sonar-scanner \
		-Dsonar.host.url=${SONARQUBE_HOST} \
		-Dsonar.login=${SONARQUBE_LOGIN} \
		-Dsonar.password=${SONARQUBE_TOKEN} \
		-Dsonar.projectKey="teler" -Dsonar.projectName="teler"