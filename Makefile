.PHONY: build clean tool lint help
all: build

# $Env:GOOS = "linux" $Env:GOOS = "darwin"
build:
	go build -o golang-gin-project-layout ./cmd/

tool:
	go vet ./...; true
	gofmt -w .

lint:
	golint ./...

clean:
	rm -rf go-gin-example
	go clean -i .

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"
