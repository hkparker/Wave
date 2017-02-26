# Wave

Wave is an 802.11 intrusion detection system, visualizer, and analytics platform.  Wireless data is sent from [collector](https://github.com/hkparker/collector)s to Wave where it is analysed by various engines.

## Developing

You'll need `npm`, `go`, and `docker-compose` available.

### Installing dependencies

Install [reflex](https://github.com/cespare/reflex), [forego](https://github.com/ddollar/forego), and [go-bindata](https://github.com/jteeuwen/go-bindata), run `go get -t` and `npm install`.

```
make deps
```

### Start instance

Start postgres and an auto-rebuilding instance of Wave.

```
make develop
```

### Running tests

Run `go test ./... -cover` and `npm test`.

```
make test
```

## Stack

* [collector](https://github.com/hkparker/collector): go application to sniff 802.11 frames and send them to Wave via websocket
* [gin](https://github.com/gin-gonic/gin): web framework in go
* [gorilla/websocket](https://github.com/gorilla/websocket): websocket support for gin
* [gorm](https://github.com/jinzhu/gorm): ORM for go used for postgres
* [postgres](https://github.com/postgres/postgres): storage of persistent data
* [react-bootstrap](https://github.com/react-bootstrap/react-bootstrap): bootstrap markup library built with react
* [vavigraphjs](https://github.com/anvaka/VivaGraphJS): graphing library with webgl support
