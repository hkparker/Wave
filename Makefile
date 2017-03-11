all: clean test release

#
# Assuming only `go` and `npm` available, install all required utilities
#
deps:
	go get github.com/cespare/reflex
	go get github.com/ddollar/forego
	go get -u github.com/jteeuwen/go-bindata/...
	make embed-assets
	go get -t
	npm install

#
# Remove all ignored files generated during build
#
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

#
# Run go-bindata to create embedded assets for single-binary deployment
#
embed-assets:
	go-bindata -pkg=helpers -o=helpers/bindata.go static/ engines/ids/rules/ engines/visualizer/metadata

#
# Download list of vendor mac prefixes from nmap
#
update-vendor-bytes:
	curl https://svn.nmap.org/nmap/nmap-mac-prefixes > engines/visualizer/metadata/nmap-mac-prefixes

#
# Run the Procfile for development
#
develop: clean embed-assets
	forego start

#
# Testing utilities
#

test-frontend:
	npm test

test-backend: embed-assets
	go test ./... -cover

test: test-frontend test-backend

#
# Building Utilities
#

build-frontend:
	./node_modules/babel-cli/bin/babel.js frontend --out-dir assets
	./node_modules/webpack/bin/webpack.js assets/* static/bundle.js

build-backend: embed-assets
	go build

build: build-frontend build-backend

#
# Build all Versions of Wave in the bin directory
#
release: clean test embed-assets update-vendor-bytes
	GOOS=linux GOARCH=amd64 go build -o bin/Wave-linux
	GOOS=linux GOARCH=arm go build -o bin/Wave-linux-arm
	GOOS=freebsd GOARCH=amd64 go build -o bin/Wave-freebsd
	GOOS=darwin GOARCH=amd64 go build -o bin/Wave-osx
	GOOS=windows GOARCH=amd64 go build -o bin/Wave-windows.exe
