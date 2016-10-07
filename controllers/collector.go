package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func pollCollector(c *gin.Context) {
	var upgrayedd websocket.Upgrader
	conn, err := upgrayedd.Upgrade(c.Writer, c.Request, nil)
	if err == nil {
		defer conn.Close()
		for {
			_, _, err = conn.ReadMessage()
			if err != nil {
				break
			}
			//fmt.Println(string(frame_bytes))
			// insert frame into IDS engine
			// insert frame into Visualizer
			// insert frame into MetadataService
		}
	}
}
