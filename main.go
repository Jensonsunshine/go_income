package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.Handle("GET", "/ping", func(context *gin.Context) {
		name := context.DefaultQuery("name", "王二狗")
		context.Writer.Write([]byte(name))
	})
	app.Run()
}
