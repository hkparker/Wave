package controllers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/hkparker/Wave/engines/visualizer"
)

var VisualEvents = make(chan map[string]string, 0)
var VisualClients = make([]*websocket.Conn, 0)

func streamVisualization(c *gin.Context) {
	var upgrayedd websocket.Upgrader
	conn, err := upgrayedd.Upgrade(c.Writer, c.Request, nil)
	if err == nil {
		defer func() {
			conn.Close()
			//delete(VisualClients, conn)
		}()
		for _, event := range visualizer.CatchupEvents() {
			err := conn.WriteJSON(event)
			if err != nil {
			}
		}
		VisualClients = append(VisualClients, conn)
	} else {
		log.WithFields(log.Fields{
			"at":    "controllers.streamVisualization",
			"error": err.Error(),
		}).Warn("failed to upgrade websocket connection")
	}
}
