CONFIG_FILE?=./config.roadie.json
PROJECT?=roadie
C = $(shell printf "\033[35;1m-->\033[0m")
V := $(if $V,,@)
GO := $(shell which go)

binaries: ; $(info $(C) building all binaries)
	$V $(MAKE) binary GOARCH=amd64 GOOS=linux
	$V $(MAKE) binary GOARCH=amd64 GOOS=windows
	$V $(MAKE) binary GOARCH=amd64 GOOS=darwin
	$V $(MAKE) binary GOARCH=386 GOOS=linux
	$V $(MAKE) binary GOARCH=386 GOOS=windows

binary: GOARCH?=amd64
binary: GOOS?=linux
binary: ; $(info $(C) building binary $(PROJECT).$(GOARCH)-$(GOOS))
	$V cp -r ./dist ./pkg/server/dist
	$V $(GO) build -o dist/$(PROJECT).$(GOARCH)-$(GOOS) ./cmd/$(PROJECT)
	$V if [ "$(GOOS)" = "windows" ]; then \
		$V mv dist/$(PROJECT).$(GOARCH)-$(GOOS) dist/$(PROJECT).$(GOARCH)-$(GOOS).exe; \
	fi

build-fe: ; $(info $(C) building the frontend assets)
	$V yarn && NODE_OPTIONS=--openssl-legacy-provider yarn build

clean: ; $(info $(C) cleaning assets and dist)
	$V rm -rf dist pkg/server/dist

coverage: ; $(info $(C) running coverage)
	$V $(GO) test -race -covermode=atomic -coverprofile=c.out ./...
	$V sed -i '' '/^github.com\/cloudcloud\/roadie\/pkg\/server\/assets.go.*/d' c.out
	$V $(GO) tool cover -html=c.out -o cover.html

# at this time, there's no watch enabled for the go binary
dev-be: bin-prep bin-dist install ; $(info $(C) building back-end for dev)
	$V CONFIG_FILE=$(CONFIG_FILE) $(PROJECT)

# dev-fe is a watch task with built-in node server
dev-fe: ; $(info $(C) building front-end for dev)
	$V NODE_OPTIONS=--openssl-legacy-provider yarn serve

docker:
	$V docker build -t cloudcloud/roadie:latest .

docker.push:
	$V docker push cloudcloud/roadie:latest

install: build-fe ; $(info $(C) installing $(PROJECT))
	$V cp -r ./dist ./pkg/server/dist
	$V $(GO) build ./cmd/$(PROJECT)/

local:
	$V GOOS=darwin $(MAKE) install
	$V HOSTNAME=http://localhost:8008 CONFIG_FILE=$(shell pwd)/config.roadie.json ./roadie

test: install ; $(info $(C) running tests)
	$V $(GO) test -v -race ./...

