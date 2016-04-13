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
			} else if user, err := ActiveSession(session_cookies[0]); err != nil {
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

func ActiveSession(id string) (user database.User, err error) {
	session := &database.Session{}
	db_err := database.DB().First(&session, "Cookie = ?", id)
	err = db_err.Error
	if err != nil {
		log.WithFields(log.Fields{
			"at":    "middleware.ActiveSession",
			"error": err.Error(),
		}).Warn("error looking up session")
		return
	}
	// ensure session hasn't expired

	db_err = database.DB().Model(&session).Related(&user)
	err = db_err.Error
	if err != nil {
		log.WithFields(log.Fields{
			"at":    "middleware.ActiveSession",
			"error": err.Error(),
		}).Warn("error finding related user for session")
	}
	return
}
