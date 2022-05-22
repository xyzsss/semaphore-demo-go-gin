package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	// 如果没有全局变量 router
	// router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", showIndexPage)
	router.Run()
}
