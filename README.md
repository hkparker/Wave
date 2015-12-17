Wave
====

Wave is a wireless Intrusion Detection System and 802.11 visualizer.  Wireless data is sent from [Wave-collector](https://github.com/hkparker/collector)s to Wave where it is analysed and stored in elasticsearch.

Goals
-----

The primary goal of Wave is to detect wireless attacks and alert relevent parties over email.  Performing wireless intrusion detection requires storing captures wireless frames for a short period of time for analysis.  During this analysis, it is also possible to construct a rich graph of live activity and relationships between devices from the wireless metadata.  Wave aims to present the information in the air as a useful and powerful graph.

Stack
-----

* [Wave-collector](https://github.com/hkparker/collector): go application to sniff 802.11 frames and send them to Wave
* [gin](https://github.com/gin-gonic/gin): a high performance web framework for web development in go
* [gorilla/websocket](https://github.com/gorilla/websocket): websocket support for sending wireless frames and frontend updates
* [gorm](https://github.com/jinzhu/gorm): ORM for go
* [ginkgo](https://github.com/onsi/ginkgo): behavior driven testing framework for go
* [bootstrap](https://github.com/twbs/bootstrap): frontend markdown toolkit
* [postgres](https://github.com/postgres/postgres): robust relational database for persistent data
* [elastalert](https://github.com/yelp/elastalert): flexible datastore for ephemeral wireless data
* [react](https://github.com/facebook/react): front end framework and router
* [d3](https://github.com/mbostock/d3): powerful javascript graphing library

Status
------

A basic [proof of concept](https://github.com/hkparker/cWave) was put together but was only a limited demonstration.  I'm currently setting up the application stucture in go.

License
-------

This project is licensed under the MIT license, see LICENSE for more information.
