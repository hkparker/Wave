package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hkparker/Wave/helpers"
	"path/filepath"
)

func EmbeddedAssets() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		switch path {
		case "/":
			RenderWebpack(c)
		default:
			data, err := helpers.Asset("static" + path)
			if err == nil {
				c.Writer.Header().Set("Content-Type", contentType(path))
				c.String(200, string(data))
				c.Abort()
			}
		}
	}
}

func RenderWebpack(c *gin.Context) {
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
	c.Abort()
}

func contentType(file string) string {
	extension := filepath.Ext(file)
	switch extension {
	case ".js":
		return "application/javascript"
	case ".svg":
		return "image/svg+xml"
	case ".ico":
		return "image/x-icon"
	case ".woff":
		return "application/font-woff"
	case ".woff2":
		return "application/font-woff"
	case ".ttf":
		return "application/font-sfnt"
	case ".eot":
		return "application/vnd.ms-fontobject"
	}
	return "application/octet-stream"
}
