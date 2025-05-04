GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

signature:
	go mod tidy
	env GO111MODULE=on go build -v $(LDFLAGS) ./cmd/signature

clean:
	rm signature

test:
	go test -v ./...

lint:
	golangci-lint run ./...

proto:
	sh ./bin/compile.sh

.PHONY: \
	signature \
	clean \
	test \
	lint \
	proto