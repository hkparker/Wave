package middleware

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/hkparker/Wave/database"
	"time"
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
				return
			}

			if len(session_cookies) != 1 {
				c.Redirect(302, "/login")
				c.Abort()
				log.WithFields(log.Fields{
					"at":     "middleware.Authentication",
					"reason": "wave_session headers does not contain exactly one string",
				}).Info("redirecting unauthenticated request")
				return
			}

			var user database.User
			if session, err := database.SessionFromID(session_cookies[0]); err == nil {
				session.LastUsed = time.Now()
				database.Orm.Save(&session)
				if user, err = session.ActiveUser(); err != nil {
					c.Redirect(302, "/login")
					c.Abort()
					log.WithFields(log.Fields{
						"at":     "middleware.Authentication",
						"reason": "could not find user for session",
					}).Info("redirecting unauthenticated request")
					return
				}
			} else {
				c.Redirect(302, "/login")
				c.Abort()
				log.WithFields(log.Fields{
					"at":     "middleware.Authentication",
					"reason": "wave_session header does not exist in session record",
				}).Info("redirecting unauthenticated request")
				return
			}

			if AdminProtected(endpoint) && !user.Admin {
				c.JSON(401, gin.H{"error": "permission denied"})
				c.Abort()
				log.WithFields(log.Fields{
					"at":             "middleware.Authentication",
					"reason":         "user is not administrator",
					"user_id":        user.ID,
					"session_cookie": session_cookies[0],
					"endpoint":       endpoint,
				}).Warn("blocking unauthenticated request")
				return
			}
		}
	}
}

//
// Given an endpoint, return if the endpoint is accessible without authentication.
//
func PublicEndpoint(url string) bool {
	switch url {
	case "/login":
		return true
	}
	return false
}

//
// Given an endpoint, return if the endpoint can only be accessed by users
// with the admin role.
//
func AdminProtected(url string) bool {
	switch url {
	case "/users/create":
		return true
	}
	return false
}
