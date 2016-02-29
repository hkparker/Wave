all:
	babel frontend --out-dir assets
	webpack assets/* static/bundle.js
	go-bindata static/
	sed -i -e 's/package main/package helpers/g' bindata.go
	mv bindata.go helpers
	go build

test:
	go test ./... -cover -race

deps:
	go get github.com/jteeuwen/go-bindata
	go install github.com/jteeuwen/go-bindata
	npm install

clean:
	rm -f Wave
	rm -f assets/*.js
	rm -f static/bundle.js
	rm -f helpers/bindata.go
