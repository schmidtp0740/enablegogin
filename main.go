package main

import (
	"enablegogin/pkg"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		pkg.Ping(c)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
