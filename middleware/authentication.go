package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hkparker/Wave/models"
	"github.com/jinzhu/gorm"
)

//
// Ensure that a request is authenticated
//
func Authentication(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !PublicEndpoint(c.Request.URL.Path) {
			session_cookie, present := c.Request.Header["wavesession"]
			if !present || len(session_cookie) != 1 {
				c.Redirect(302, "/login")
				c.Abort()
			} else if !ActiveSession(session_cookie[0], db) {
				c.Redirect(302, "/login")
				c.Abort()
			} else if !HasRole() {
				// JSON permission denied
			}
		}
	}
}

//
// All endpoints that can be accessed without
// authentication need to be explicitly called
// out here.
//
func PublicEndpoint(url string) bool {
	switch url {
	case "/frames":
		return true
	case "/login":
		return true
	}
	return false
}

//
// See if the provided session cookie is valid in the database.
//
func ActiveSession(id string, db *gorm.DB) bool {
	//Session.where(id: id).nil?
	session := &models.Session{}
	db.Where(&models.Session{Cookie: id}).First(&session)
	return true
}

//
// Ensure the user associated with the session has
// the role required for the endpoint.
//
func HasRole() bool {
	return true
}
