//go:build !excludetest

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/weidaicheng/github-action-sharing/service-greeting/svc"
)

func main() {
	r := gin.Default()

	r.GET("/greeting/:name", func(c *gin.Context) {
		name := c.Param("name")
		message := svc.SayHello(name)

		c.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	})

	r.Run()
}
