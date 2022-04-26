-include .env

VERSION := $(shell git describe --tags --always)
BUILD := $(shell git rev-parse --short HEAD)
LDFLAGS= -ldflags "-X main.Version=$(VERSION) -X main.Build=$(BUILD)"

build:
	@echo "  >  Building binary..." \
	go mod tidy
	go build  -o ./build/ $(LDFLAGS) ./cmd/ethos && \
	cd ./cmd/ethos && go install $(LDFLAGS)

clean:
	@echo "  >  Cleaning artifacts..." \
	$(RM) build

