.DEFAULT_GOAL := build
.PHONY := build

BINDIR       := $(CURDIR)/.bin
BINNAME      ?= proxx

build:
	go get ./cmd/$(BINNAME) && GO111MODULE=on CGO_ENABLED=0 go build -o $(BINDIR)/$(BINNAME) ./cmd/$(BINNAME)
