package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.Handle("GET", "/ping", func(c *gin.Context) {
		name := c.DefaultQuery("name", "ηδΊη")
		c.Writer.Write([]byte(name))
	})
	app.Run(":888")
}
