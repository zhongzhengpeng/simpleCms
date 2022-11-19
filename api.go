package main

import (
	"simpleCms/app/content"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func init() {
	router.GET("/content/:id", content.GET)
	router.POST("/content", content.POST)
	router.PUT("/content/:id", content.PUT)
	router.DELETE("/content/:id", content.DELETE)
}
