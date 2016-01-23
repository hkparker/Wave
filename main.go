package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/hkparker/Wave/controllers"
)

func renderWebpack(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/html")
	c.String(200,
		`<html>
	<head>
		<meta charset="utf-8">
		<title>Wave</title>
	</head>
	<body>
		<div id="content"></div>
		<script type="text/javascript" src="bundle.js" charset="utf-8"></script>
	</body>
</html>
`,
	)
}

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(static.Serve("/vendor", static.LocalFile("vendor", false)))
	router.Use(static.Serve("/", static.LocalFile("static", false)))
	router.GET("/", renderWebpack)
	router.GET("/frames", controllers.PollCollector)
	return router
}

func main() {
	//gin.SetMode(gin.ReleaseMode)
	NewRouter().Run(":8080")
}
