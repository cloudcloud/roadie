CONFIG_FILE?=./config.roadie.json
PROJECT?=roadie
C = $(shell printf "\033[35;1m-->\033[0m")
V := $(if $V,,@)

bin-prep: ; $(info $(C) preparing for bin-data)
	$V GO111MODULE=off go get -u github.com/kevinburke/go-bindata/...

bin-dist: ; $(info $(C) collapsing files into bin-data)
	$V go-bindata -o ./pkg/server/assets.go -pkg server -prefix dist/ dist/...

binaries: ; $(info $(C) building all binaries)
	$V $(MAKE) binary GOARCH=amd64 GOOS=linux
	$V $(MAKE) binary GOARCH=amd64 GOOS=windows
	$V $(MAKE) binary GOARCH=amd64 GOOS=darwin
	$V $(MAKE) binary GOARCH=386 GOOS=linux
	$V $(MAKE) binary GOARCH=386 GOOS=windows

binary: GOARCH?=amd64
binary: GOOS?=linux
binary: ; $(info $(C) building binary $(PROJECT).$(GOARCH)-$(GOOS))
	$V go build -o dist/$(PROJECT).$(GOARCH)-$(GOOS) ./cmd/$(PROJECT)
	$V if [ "$(GOOS)" = "windows" ]; then \
		$V mv dist/$(PROJECT).$(GOARCH)-$(GOOS) dist/$(PROJECT).$(GOARCH)-$(GOOS).exe; \
	fi

build-fe: ; $(info $(C) building the frontend assets)
	$V yarn && NODE_OPTIONS=--openssl-legacy-provider yarn build

clean: ; $(info $(C) cleaning assets and dist)
	$V rm pkg/server/assets.go
	$V rm -r dist

coverage: ; $(info $(C) running coverage)
	go test -race -covermode=atomic -coverprofile=c.out ./...
	sed -i '' '/^github.com\/cloudcloud\/roadie\/pkg\/server\/assets.go.*/d' c.out
	go tool cover -html=c.out -o cover.html

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

install: build-fe bin-dist ; $(info $(C) installing $(PROJECT))
	$V go build ./cmd/$(PROJECT)/

local:
	$V GOOS=darwin $(MAKE) install
	$V HOSTNAME=http://localhost:8008 CONFIG_FILE=$(shell pwd)/config.roadie.json ./roadie

test: install ; $(info $(C) running tests)
	$V go test -v -race ./...

