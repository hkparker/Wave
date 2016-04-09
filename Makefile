all: clean test build

clean:
	@rm -f Wave
	@rm -f assets/*.js
	@rm -f static/bundle.js
	@rm -f static/448c34a56d699c29117adc64c43affeb.woff2
	@rm -f static/89889688147bd7575d6327160d64e760.svg
	@rm -f static/e18bbf611f2a2e43afc071aa2f4e1512.ttf
	@rm -f static/f4769f9bdb7466be65088239c12046d1.eot
	@rm -f static/fa2772327f55d8198301fdb8bcfc8158.woff
	@rm -f helpers/bindata.go
	@rm -f bin/*

embed-assets:
	go-bindata -pkg=helpers -o=helpers/bindata.go static/

develop: clean build
	WAVE_ENV=development ./Wave
	# hot reloading react, redux dev tools, etc
	# rebuild and restart Wave when go files change

test-frontend:

test-backend: embed-assets
	go test ./... -cover

test: test-frontend test-backend

build-frontend:
	babel frontend --out-dir assets
	webpack assets/* static/bundle.js

build-backend: embed-assets
	go build

build: build-frontend build-backend

release: clean test
	GOOS=linux GOARCH=amd64 go build -o bin/Wave-linux
	GOOS=linux GOARCH=arm go build -o bin/Wave-linux-arm
	GOOS=freebsd GOARCH=amd64 go build -o bin/Wave-freebsd
	GOOS=darwin GOARCH=amd64 go build -o bin/Wave-da
	GOOS=windows GOARCH=amd64 go build -o bin/Wave-windows.exe
