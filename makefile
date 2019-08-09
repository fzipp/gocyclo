PLATFORM=$(shell uname -s | tr '[:upper:]' '[:lower:]')

.PHONY: build clean dist

build:
	go fmt ./...
	@mkdir -p ./bin/
	CGO_ENABLED=1 go build -o ./bin/gocyclo github.com/adamdecaf/gocyclo

clean:
	@rm -rf ./bin/

dist: clean build
ifeq ($(OS),Windows_NT)
	CGO_ENABLED=1 GOOS=windows go build -o bin/gocyclo-windows-amd64.exe github.com/adamdecaf/gocyclo
else
	CGO_ENABLED=1 GOOS=$(PLATFORM) go build -o bin/gocyclo-$(PLATFORM)-amd64 github.com/adamdecaf/gocyclo
endif
