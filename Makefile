PROJECT_ROOT?=$(shell pwd)
GO_BINDATA_BIN=$(VENDOR_DIR)/github.com/jteeuwen/go-bindata/go-bindata/go-bindata

.PHONY: build 
build: generate
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o .out/gofoody-app

.PHONY: generate
generate:
	go-bindata -pkg templates -o templates/bindata.go -nocompress -ignore=.*go templates/...

.PHONY: clean
clean:
	-rm -rf .out/
	-rm -f ./templates/bindata.go

.PHONY: build-ci
build-ci: build
	docker build -f Dockerfile -t gofoody-app:latest .
