all:
	babel frontend --out-dir assets
	webpack assets/* static/bundle.js
	go build

backend-tests:
	go test ./... -cover -race

frontend-tests:
	echo "frontend tests missing!"

test: backend-tests frontend-tests

setup-env:
	npm install

clean:
	rm -f Wave
	rm -f assets/*.js
	rm -f static/bundle.js
