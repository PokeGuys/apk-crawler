# Change these variables as necessary.
MAIN_PACKAGE_PATH := ./cmd/crawler/
TMP_DIR := ./tmp
BINARY_NAME := apk-crawler

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

.PHONY: prepare
prepare:
	mkdir -p ${TMP_DIR}/bin
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
	go install github.com/vektra/mockery/v2@v2.44.1
	go env -w CGO_ENABLED=1

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

.PHONY: no-dirty
no-dirty:
	git diff --exit-code


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit:
	go mod verify
	go vet ./...
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.1 run ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...


# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

.PHONY: proto
proto:
	@protoc --go_out=. --go_opt=paths=source_relative ./proto/*.proto

.PHONY: mocks
mocks:
	@mockery --all

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=${TMP_DIR}/coverage.out.tmp ./...
	cat ${TMP_DIR}/coverage.out.tmp | grep -Ev "mock_|mocks/|cmd/|.pb.go" > ${TMP_DIR}/coverage.out
	go tool cover -html=${TMP_DIR}/coverage.out

## build: build the application
.PHONY: build
build:
	go build -o=${TMP_DIR}/bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

## run: run the application
.PHONY: run
run: build
	${TMP_DIR}/bin/${BINARY_NAME}

## run/live: run the application with reloading on file changes
.PHONY: run/live
run/live:
	go run github.com/cosmtrek/air@v1.51.0


# ==================================================================================== #
# OPERATIONS
# ==================================================================================== #

## push: push changes to the remote Git repository
.PHONY: push
push: tidy audit no-dirty
	git push

## production/deploy: deploy the application to production
.PHONY: production/build
production/build:
	GOOS=linux GOARCH=amd64 go build -ldflags='-s -w' -o=${TMP_DIR}/bin/linux_amd64/${BINARY_NAME} ${MAIN_PACKAGE_PATH}
