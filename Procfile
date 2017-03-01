wave-api: reflex -r '.*.go$' -s -- sh -c 'go run main.go'
wave-rules: reflex -r 'engines/ids/rules/.*.js$' -s -- sh -c 'make embed-assets'
#wave-frontend: reflex -r 'frontend/.+\.js.?$' -s -- sh -c 'make build-frontend'
wave-services: docker-compose up
