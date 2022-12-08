package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.Handle("GET", "/ping", func(c *gin.Context) {
		name := c.DefaultQuery("name", "王二狗")
		c.Writer.Write([]byte(name))
	})
	app.Run()
}
