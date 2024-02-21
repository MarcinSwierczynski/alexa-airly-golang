.DEFAULT_GOAL := build

.PHONY: fmt vet build package clean
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap *.go

package: build
	zip alexa-airly-golang bootstrap

clean:
	go clean