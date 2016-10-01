package controllers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
)

func newSession(c *gin.Context) {}

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
