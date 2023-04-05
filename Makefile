PKGNAME=seeddms-munin
VERSION=0.0.1
TARGET=http://seeddms.steinmann.cx/restapi/index.php
APIKEY=d34c8dab1483f51a9f6a0d8dc8d5348c

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet
GOGET=$(GOCMD) get

default: build

build: test
	mkdir -p bin
	$(GOBUILD) -v -o bin/${PKGNAME} main.go

clean:
	$(GOCLEAN)
	rm -rf bin

test:
	$(GOVET) ./...
	$(GOTEST) -v ./...

run:
	target=${TARGET} apikey=${APIKEY} bin/${PKGNAME} run

dist: clean
	rm -rf ${PKGNAME}-${VERSION}
	mkdir ${PKGNAME}-${VERSION}
	cp -r *.go Makefile go.mod ${PKGNAME}-${VERSION}
	tar czvf ${PKGNAME}-${VERSION}.tar.gz ${PKGNAME}-${VERSION}
	rm -rf ${PKGNAME}-${VERSION}

.PHONY: build test 
