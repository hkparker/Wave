package controllers

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

//frame_elasticsearch_client
var transport = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	Dial: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: time.Minute,
	}).Dial,
	TLSHandshakeTimeout: 10 * time.Second,
}

var client = &http.Client{
	Transport: transport,
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// make sure the client presents proper TLS client cert
		return true
	},
}

func elasticache(frame []byte) {
	req, err := http.NewRequest("POST", "http://127.0.0.1:9200/frames/frame/", bytes.NewBuffer(frame))
	if err != nil {
		fmt.Println("error creating elasticsearch request:", err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error writing frame to elasticsearch:", err)
		return
	}
	ioutil.ReadAll(resp.Body)
	// parse ID, go delete in 30 seconds
	//go func() {
	//	time.Sleep(30 * time.Second)
	//	make a delete request for that frame
	//	client.Do(reqQ)
	//}()
	resp.Body.Close()
}

func PollCollector(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err == nil {
		defer conn.Close()
		for {
			_, frame_bytes, err := conn.ReadMessage()
			if err != nil {
				break
			}
			//frame := string(frame_bytes)
			elasticache(frame_bytes)
			// update visualizer
		}
	}
}
