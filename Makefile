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
	go get golang.org/x/crypto/bcrypt
	go get gopkg.in/olivere/elastic.v3
	go get github.com/jinzhu/gorm
	go get github.com/sec51/twofactor
	go get github.com/stretchr/testify
	go get github.com/gin-gonic/gin
	go get github.com/gorilla/websocket
	go get github.com/Sirupsen/logrus
	go get github.com/lib/pq
	go get github.com/jteeuwen/go-bindata
	go install github.com/jteeuwen/go-bindata
	npm install

clean:
	rm -f Wave
	rm -f assets/*.js
	rm -f static/bundle.js
	rm -f static/448c34a56d699c29117adc64c43affeb.woff2
	rm -f static/89889688147bd7575d6327160d64e760.svg
	rm -f static/e18bbf611f2a2e43afc071aa2f4e1512.ttf
	rm -f static/f4769f9bdb7466be65088239c12046d1.eot
	rm -f static/fa2772327f55d8198301fdb8bcfc8158.woff
	rm -f helpers/bindata.go
