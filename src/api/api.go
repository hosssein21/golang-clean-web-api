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

	v1:= r.Group("/api/v1/")
	{
		// v1.GET("/health", func(c *gin.Context){
		// 	c.JSON(http.StatusOK,"its working!")
		// 	return
		// })

		health:=v1.Group("/health")
		routers.FirstRouter(health)
	}

	r.Run(fmt.Sprintf(":%s",cfg.Server.InternalPort))

}