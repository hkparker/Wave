package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/hkparker/Wave/engines/visualizer"
	"github.com/hkparker/Wave/models"
	log "github.com/sirupsen/logrus"
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

func currentUser(c *gin.Context) (models.User, error) {
	userInterface, ok := c.Get("currentUser")
	if !ok {
		c.Status(500)
		c.Abort()
		errStr := "request context has no currentUser set"
		log.WithFields(log.Fields{
			"at": "controllers.currentUser",
		}).Error(errStr)
		return models.User{}, errors.New(errStr)
	}
	user, ok := userInterface.(models.User)
	if !ok {
		c.Status(500)
		c.Abort()
		errStr := "unable to cast user interface to user model"
		log.WithFields(log.Fields{
			"at": "controllers.currentUser",
		}).Error(errStr)
		return user, errors.New(errStr)
	}

	return user, nil
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
