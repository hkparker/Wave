wave-api: reflex -r '\.go$' -s -- sh -c 'go build . && WAVE_ENV=development ./Wave --port=8080 --collector-port=4444  --db_username=postgres  --db_password=postgres'
#wave-frontend: reflex -r 'frontend/.+\.js.?$' -s -- sh -c 'make build-frontend'
wave-services: docker-compose up
