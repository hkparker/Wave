package middleware

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/hkparker/Wave/database"
)

//
// Ensure that a request is authenticated
//
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		endpoint := c.Request.URL.Path
		if !PublicEndpoint(endpoint) {
			session_cookies, present := c.Request.Header["wave_session"]
			if !present {
				c.Redirect(302, "/login")
				c.Abort()
				log.WithFields(log.Fields{
					"at":     "middleware.Authentication",
					"reason": "missing wave_session header",
				}).Info("redirecting unauthenticated request")
			} else if len(session_cookies) != 1 {
				c.Redirect(302, "/login")
				c.Abort()
				log.WithFields(log.Fields{
					"at":     "middleware.Authentication",
					"reason": "wave_session headers does not contain exactly one string",
				}).Info("redirecting unauthenticated request")
			} else if active, user := ActiveSession(session_cookies[0]); !active {
				c.Redirect(302, "/login")
				c.Abort()
				log.WithFields(log.Fields{
					"at":     "middleware.Authentication",
					"reason": "wave_session header is not an active session",
				}).Info("redirecting unauthenticated request")
			} else if AdminProtected(endpoint) && !user.Admin {
				c.JSON(401, gin.H{"error": "permission denied"})
				log.WithFields(log.Fields{
					"at":             "middleware.Authentication",
					"reason":         "user is not administrator",
					"user_id":        user.ID,
					"session_cookie": session_cookies[0],
				}).Warn("blocking unauthenticated request")
			}
		}
	}
}

//
// All endpoints that can be accessed without authentication need to be
// explicitly called out here.  Static content
//
func PublicEndpoint(url string) bool {
	switch url {
	case "/login":
		return true
	}
	return false
}

func AdminProtected(url string) bool {
	switch url {
	case "/users/create":
		return true
	}
	return false
}

func ActiveSession(id string) (bool, database.User) {
	//Session.where(id: id).nil?
	//session := &models.Session{}
	//db.Where(&models.Session{Cookie: id}).First(&session)
	return true, database.User{Admin: true}
}
