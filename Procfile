wave-api: reflex -r '\.go$' -s -- sh -c 'go build . && WAVE_ENV=development ./Wave --port=8080  --db_username=postgres  --db_password=postgres'
wave-frontend: reflex -r 'frontend/.+\.js.?$' -s -- sh -c 'babel frontend --out-dir assets && webpack assets/* static/bundle.js'
wave-services: docker-compose up
