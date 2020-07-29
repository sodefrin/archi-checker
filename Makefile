PACKAGES = $(shell go list ./... )

.PHONY: build
build:
	go build ./main.go

.PHONY: test
test:
	go test -v $(PACKAGES)
