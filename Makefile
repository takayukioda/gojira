PACKAGE=$(shell go list .)
FILES=./...

.PHONY: build fmt run

build:
	env GO111MODULE=on go build $(FILES)
fmt:
	go fmt $(FILES)
run:
	go run $(FILES)
