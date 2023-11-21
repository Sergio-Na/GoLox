# Makefile for building and running a Go application

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
BINARY_NAME=goLox

all: build

build: 
	$(GOBUILD) -o $(BINARY_NAME) .

run:
	$(GORUN) main.go

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

.PHONY: all build run clean
