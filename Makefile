all:
	go build -o Wave main.go

backend-tests:
	go test ./... -cover -race

test: backend-tests

clean:
	rm Wave
	rm -r assets/*
	rm -r static/*
