package router

import (
	"github.com/gin-gonic/gin"
	midware "secret/utils/middle"
)

func Router() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.Recovery())
	//router.Use(midware.Cors())
	router.Use(midware.NoRoute())
	router.Use(midware.Limit)
	router.Use(midware.Recover("baiy"))
	rf := router.Group("/api/v1")
	UserRouter(rf)
	return router
}
