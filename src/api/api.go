package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hosssein21/golang-clean-web-api/api/routers"
	"github.com/hosssein21/golang-clean-web-api/config"
)

func InitServer() {
	cfg:=config.GetConfig()
	r := gin.New()
	r.Use(gin.Logger(),gin.Recovery())

	api:= r.Group("/api")

	v1:=api.Group("/v1")
	{
		// v1.GET("/health", func(c *gin.Context){
		// 	c.JSON(http.StatusOK,"its working!")
		// 	return
		// })

		health:=v1.Group("/health")
		routers.FirstRouter(health)

		test_routers:=v1.Group("/test")
		routers.TestRouters(test_routers)
	}

	r.Run(fmt.Sprintf(":%s",cfg.Server.InternalPort))

}