package router

import (
	"github.com/Niexiawei/golang-skeleton/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Cors)
	return r
}
