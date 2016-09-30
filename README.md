# Wave

Wave is a wireless Intrusion Detection System and 802.11 visualizer.  Wireless data is sent from [collector](https://github.com/hkparker/collector)s to Wave where it is analysed and stored in elasticsearch.  In early development.

## Developing

You'll need `npm`, `go`, and `docker-compose` available.

### Installing dependencies

Install reflex, forego, and go-bindata, run `go-get -t` and `npm install`.

```
make deps
```

### Start instance

Run `make develop` to start elasticsearch and postgres as well as an auto-rebuilding instance of Wave.

### Running tests

```
make test
```

## Stack

* [collector](https://github.com/hkparker/collector): go application to sniff 802.11 frames and send them to Wave via websocket
* [gin](https://github.com/gin-gonic/gin): web framework in go
* [gorilla/websocket](https://github.com/gorilla/websocket): websocket support for gin
* [gorm](https://github.com/jinzhu/gorm): ORM for go used for postgres
* [postgres](https://github.com/postgres/postgres): storage of persistent data
* [elasticsearch](https://github.com/elastic/elasticsearch): storage of ephemeral wireless frames
* [elastigo](https://github.com/mattbaird/elastigo): go client for elasticsearch
* [react-bootstrap](https://github.com/react-bootstrap/react-bootstrap): bootstrap markup library built with react
* [vavigraphjs](https://github.com/anvaka/VivaGraphJS): graphing library with webgl support
