# Set an output prefix, which is the local directory if not specified
PREFIX?=$(shell pwd)
# Setup name variables for the package/tool
NAME := k8s-svc-endpoint
PKG := github.com/codetime66/$(NAME)
# Set our default go compiler
GO := go
#
VERSION := $(shell cat VERSION.txt)

.PHONY: build
build: $(NAME) 

$(NAME): $(wildcard *.go) $(wildcard */*.go) VERSION.txt
	@echo "+ $@"
	$(GO) build -o ./bin/k8s-svc-endpoint ./cmd/k8s-svc-endpoint
