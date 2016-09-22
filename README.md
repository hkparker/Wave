# Wave

Wave is a wireless Intrusion Detection System and 802.11 visualizer.  Wireless data is sent from [collector](https://github.com/hkparker/collector)s to Wave where it is analysed and stored in elasticsearch.  In early development.

## Goals

The primary goal of Wave is to detect wireless attacks and alert relevent parties.  Performing wireless intrusion detection requires storing captures wireless frames for a short period of time for analysis.  During this analysis, it is also possible to construct a rich graph of live activity and relationships between devices from the wireless metadata.  Wave aims to present the information in the air as a powerful interactive graph.

## Developing

### Installing dependencies

Requires only `npm` and `go`

```
make deps
```

### Start instance

Run `make develop` to start elasticsearch and postgres as well as an auto restarting instance of Wave.

### Running tests

```
make test
```

## Deploying

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

## License

This project is licensed under the MIT license, see LICENSE for more information.
