GO = go
GOTEST ?= go test
export TEST_COUNT ?= 1
export TEST_ARGS ?=
bin:
	mkdir -p bin/

.PHONY: build
build: bin
	$(GO) build  -ldflags "${LDFLAGS}" -o bin/coinconv ./cmd/coinconv/*.go

clear:
	rm -rf bin

.PHONY: test
test:
	$(GO) test -v ./... -count $(TEST_COUNT) -race $(TEST_ARGS)