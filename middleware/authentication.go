package middleware

import (
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !PublicEndpoint(c.Request.URL.Path) {
			if !ActiveSession() {
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
	case "/":
		return true
	case "/bundle.js":
		return true
	case "/wave.svg":
		return true
	case "/frames":
		return true
	}
	return false
}

//
// See if the provided session cookie is valid in
// the database.
//
func ActiveSession() bool {
	return true
}

//
// Ensure the user associated with the session has
// the role required for the endpoint.
//
func HasRole() bool {
	return true
}
