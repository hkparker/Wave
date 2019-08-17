package controllers

import (
	log "github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/hkparker/Wave/engines/visualizer"
	"github.com/hkparker/Wave/models"
	"net/http"
)

func init() {
	go func() {
		for {
			event := <-visualizer.VisualEvents
			VisualClientMux.Lock()
			clients_list := VisualClients
			VisualClientMux.Unlock()
			for id, client := range clients_list {
				err := client.WriteJSON(event)
				if err != nil {
					log.WithFields(log.Fields{
						"at":    "controllers.init",
						"error": err.Error(),
					}).Info("visualizer client disconnected")
					VisualClientMux.Lock()
					delete(VisualClients, id)
					VisualClientMux.Unlock()
				}
			}
		}
	}()
}

func requestJSON(c *gin.Context) (json map[string]string, err error) {
	err = c.BindJSON(&json)
	if err != nil {
		c.JSON(400, gin.H{"error": "error parsing json"})
		c.Abort()
		log.WithFields(log.Fields{
			"at":    "controllers.requestJSON",
			"error": err.Error(),
		}).Error("error parsing request")
	}
	return
}

func currentUser(c *gin.Context) (user models.User, err error) {
	session_cookie, err := sessionCookie(c)
	if err != nil {
		return
	}
	user, err = models.UserFromSessionCookie(session_cookie)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		c.Abort()
		log.WithFields(log.Fields{
			"at":    "controllers.currentUser",
			"error": err.Error(),
		}).Error("error getting current user")
	}
	return
}

func sessionCookie(c *gin.Context) (session_cookie string, err error) {
	var session_cookie_obj *http.Cookie
	session_cookie_obj, err = c.Request.Cookie("wave_session")
	if err != nil {
		log.WithFields(log.Fields{
			"at":    "controllers.sessionCookie",
			"error": err.Error(),
		}).Error("session_cookie_missing")
		return
	} else {
		session_cookie = session_cookie_obj.Value
	}
	return
}
