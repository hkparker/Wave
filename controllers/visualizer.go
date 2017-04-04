package controllers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/hkparker/Wave/engines/visualizer"
	"github.com/satori/go.uuid"
	"net/http"
	"sync"
)

var VisualClients = make(map[string]*websocket.Conn, 0)
var VisualClientMux sync.Mutex

func streamVisualization(c *gin.Context) {
	upgrayedd := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrayedd.Upgrade(c.Writer, c.Request, nil)
	if err == nil {
		id := uuid.NewV4().String()
		for _, event := range visualizer.CatchupEvents() {
			if len(event) > 0 {
				err := conn.WriteJSON(event)
				if err != nil {
					log.WithFields(log.Fields{
						"at":    "controllers.streamVisualization",
						"error": err.Error(),
					}).Error("error writing catch-up event")
				}
			}
		}
		VisualClientMux.Lock()
		VisualClients[id] = conn
		VisualClientMux.Unlock()
	} else {
		log.WithFields(log.Fields{
			"at":    "controllers.streamVisualization",
			"error": err.Error(),
		}).Warn("failed to upgrade websocket connection")
	}
}
