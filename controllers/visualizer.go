package controllers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/hkparker/Wave/engines/visualizer"
	"github.com/satori/go.uuid"
)

var VisualEvents = make(chan map[string]string, 0)
var VisualClients = make(map[string]*websocket.Conn, 0)

func streamVisualization(c *gin.Context) {
	var upgrayedd websocket.Upgrader
	conn, err := upgrayedd.Upgrade(c.Writer, c.Request, nil)
	if err == nil {
		id := uuid.NewV4().String()
		defer func() {
			conn.Close()
			if _, present := VisualClients[id]; present {
				delete(VisualClients, id)
			}
		}()
		for _, event := range visualizer.CatchupEvents() {
			err := conn.WriteJSON(event)
			if err != nil {
			}
		}
		VisualClients[id] = conn
	} else {
		log.WithFields(log.Fields{
			"at":    "controllers.streamVisualization",
			"error": err.Error(),
		}).Warn("failed to upgrade websocket connection")
	}
}
