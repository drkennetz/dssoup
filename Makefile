makeFileDir := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

test:
	go test -v -cover ./...

.PHONY: test