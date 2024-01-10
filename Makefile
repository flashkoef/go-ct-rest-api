BINARY_NAME=app
.DEFAULT_GOAL := run

# ----------------
# helpers
# ----------------

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

# ----------------
# qualitiy control
# ----------------

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
audit:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go test -race -buildvcs -vet=off ./...

# ----------------
# development
# ----------------

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test ./... -coverprofile=coverage.out

## build: build the application
.PHONY: build
build:
	GOARCH=amd64 GOOS=windows go build -o ./target/${BINARY_NAME}-windows main.go
	GOARCH=amd64 GOOS=linux go build -o ./target/${BINARY_NAME}-linux main.go

## run: run the  application
.PHONY: run
run: build
	./target/${BINARY_NAME}-windows

## run/locale run main.go locale
.PHONY: run/locale
run/locale:
	go run main.go

# ----------------
# operations
# ----------------

## push: push changes to the remote Git repository
.PHONY: push
push: tidy audit no-dirty
	git push

## clean: removes object files from package source directories and binary files in target
.PHONY: clean
clean:
	go clean
	rm ./target/${BINARY_NAME}-windows
	rm ./target/${BINARY_NAME}-linux

## dep: dep downloads the named modules into the module cache
.PHONY: dep
dep:
	go mod download

## vet: vet examines Go source code and reports suspicious constructs
.PHONY: vet
vet:
	go vet

## lint: lint starts the linter used
.PHONY: lint
lint:
	golangci-lint run --enable-all
