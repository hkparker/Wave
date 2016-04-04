# Wave

Wave is a wireless Intrusion Detection System and 802.11 visualizer.  Wireless data is sent from [collector](https://github.com/hkparker/collector)s to Wave where it is analysed and stored in elasticsearch.

## Goals

The primary goal of Wave is to detect wireless attacks and alert relevent parties.  Performing wireless intrusion detection requires storing captures wireless frames for a short period of time for analysis.  During this analysis, it is also possible to construct a rich graph of live activity and relationships between devices from the wireless metadata.  Wave aims to present the information in the air as a powerful interactive graph.

## Installation

# Building

`make`

## Developing

### Installing dependencies

Use NPM and Go to bring in dependencies

```
go get
npm install
```

### Setting up development services

Use docker-compose to start a postgres and elasticsearch servers for development and testing.

`docker-compose up`

### Start development instance

In another terminal start the development environment.

`make develop`

### Running tests

`make test`

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

## Structure

```
/assets				styles, libraries, and compiled javascript for webpack
/bin				self-contained binaries for Linux, OSX, and Windows, present on release branches
/controllers		controller logic
/database			logic for connecting to databases
/frontend			jsx source, consumed by babel and saved in assets
/helpers
/ids				logic for intrusion detection
LICENSE				the MIT license
main.go				application setup and routing
package.json		npm project file
README.md			this file
/visualizer			logic for updating frontend visualization
```

## License

This project is licensed under the MIT license, see LICENSE for more information.
