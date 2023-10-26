package api

import (

	"github.com/gin-gonic/gin"
	"github.com/hosssein21/golang-clean-web-api/api/routers"
)

func InitServer() {
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

	r.Run(":5005")

}