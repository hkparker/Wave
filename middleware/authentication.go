package middleware

import (
	log "github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/hkparker/Wave/models"
)

var public_endpoints map[string]bool
var admin_endpoints map[string]bool

func init() {
	public_endpoints = map[string]bool{
		"/login":              true,
		"/sessions/create":    true,
		"/bundle.js":          true,
		"/streams/visualizer": true,
		"/ids":                true,
	}

	admin_endpoints = map[string]bool{
		"/users":                 true,
		"/users/create":          true,
		"/users/delete":          true,
		"/users/assign-password": true,
		"/admin/tls":             true,
		"/collectors":            true,
		"/collectors/create":     true,
		"/collectors/delete":     true,
	}
}

//
// Ensure that a request is authenticated
//
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		endpoint := c.Request.URL.Path
		if _, public := public_endpoints[endpoint]; !public {
			session_cookie, err := c.Request.Cookie("wave_session")
			if err != nil {
				c.Redirect(302, "/login")
				c.Abort()
				log.WithFields(log.Fields{
					"at":       "middleware.Authentication",
					"reason":   "missing wave_session cookie",
					"error":    err.Error(),
					"endpoint": endpoint,
				}).Info("redirecting unauthenticated request")
				return
			}

			var user models.User
			if session, err := models.SessionFromID(session_cookie.Value); err == nil {
				if user, err = session.User(); err != nil {
					c.Status(401)
					c.Abort()
					log.WithFields(log.Fields{
						"at":       "middleware.Authentication",
						"reason":   "could not find user for session",
						"endpoint": endpoint,
					}).Info("redirecting unauthenticated request")
					return
				}
			} else {
				c.Status(401)
				c.Abort()
				log.WithFields(log.Fields{
					"at":       "middleware.Authentication",
					"reason":   "wave_session header does not exist in session record",
					"endpoint": endpoint,
				}).Info("redirecting unauthenticated request")
				return
			}

			c.Set("currentUser", user)

			if _, admin_protected := admin_endpoints[endpoint]; admin_protected && !user.Admin {
				c.JSON(401, gin.H{"error": "permission denied"})
				c.Abort()
				log.WithFields(log.Fields{
					"at":       "middleware.Authentication",
					"reason":   "user is not administrator",
					"user_id":  user.ID,
					"endpoint": endpoint,
				}).Warn("blocking unauthenticated request")
				return
			}
		}
	}
}
