all:
	babel frontend --out-dir assets
	webpack assets/* static/bundle.js
	go build

test:
	go test ./... -cover -race

deps:
	npm install

clean:
	rm -f Wave
	rm -f assets/*.js
	rm -f static/bundle.js
