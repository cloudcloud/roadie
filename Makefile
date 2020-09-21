PROJECT?=roadie

bin-prep:
	GO111MODULE=off go get -u github.com/kevinburke/go-bindata/...

bin-migrations:
	go-bindata -o data/migrate.go -prefix data/migrations/ data/migrations/
	sed -i "s/package main/package data/" data/migrate.go

bin-dist:
	go-bindata -o $(PROJECT)/assets.go -prefix dist/ dist/...

binaries:
	$(MAKE) binary GOARCH=amd64 GOOS=linux
	$(MAKE) binary GOARCH=amd64 GOOS=windows
	$(MAKE) binary GOARCH=amd64 GOOS=darwin
	$(MAKE) binary GOARCH=386 GOOS=linux
	$(MAKE) binary GOARCH=386 GOOS=windows

binary: GOARCH?=amd64
binary: GOOS?=linux
binary:
	go build -o build/$(PROJECT).$(GOARCH)-$(GOOS) ./$(PROJECT)
	if [ "$(GOOS)" = "windows" ]; then \
		mv build/$(PROJECT).$(GOARCH)-$(GOOS) build/$(PROJECT).$(GOARCH)-$(GOOS).exe; \
	fi

build-fe:
	yarn && yarn build

clean:
	rm $(PROJECT)/assets.go
	rm data/migrate.go
	rm -r dist

# at this time, there's no watch enabled for the go binary
dev-be: bin-prep bin-migrations bin-dist install
	$(PROJECT)

# serve is a watch task with built-in node server
dev-fe:
	yarn serve

install: build-fe bin-migrations bin-dist
	go install ./$(PROJECT)/

test: install
	go test -v -race ./...

