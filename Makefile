.PHONY: build
build:
	go build -v ./cmd/main.go

clean:
	del main.exe

.DEFAULT_GOAL := build