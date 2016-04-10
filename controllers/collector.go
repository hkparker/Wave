package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// make sure the client presents proper TLS client cert
		return true
	},
}

func PollCollector(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err == nil {
		defer conn.Close()
		for {
			_, _, err = conn.ReadMessage()
			if err != nil {
				break
			}
			//frame := string(frame_bytes)
			//database.ElasticacheFrame(frame_bytes)
			// update visualizer
		}
	}
}
