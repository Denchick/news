.PHONY: build
build:
	go build -v ./cmd/apiserver
	go build -v ./cmd/rssparser

.DEFAULT_GOAL := build