BINARY_NAME=wtools
HASH := $(shell git rev-parse HEAD)
COMMIT_DATE := $(shell git show -s --format=%ci ${HASH})
BUILD_DATE := $(shell date '+%Y-%m-%d %H:%M:%S')
LONG_HASH := ${HASH} (${COMMIT_DATE})
VERSION := "v0.0.1"

.PHONY: build
build:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME} -ldflags="-X 'github.com/krippz/wtools/cmd.version=${VERSION}' -X 'github.com/krippz/wtools/cmd.gitCommit=${LONG_HASH}' -X 'github.com/krippz/wtools/cmd.buildDate=${BUILD_DATE}'"  main.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}.exe main.go

test:
	go test ./internal/jwt -v

run:
	./${BINARY_NAME}

build_and_run: build run

clean:
	go clean