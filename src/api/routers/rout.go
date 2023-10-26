package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hosssein21/golang-clean-web-api/api/handlers"
)

func FirstRouter(r *gin.RouterGroup) {
	handler:= handlers.NewHealthHandler()

	r.GET("1/",handler.Health)
	r.POST("/",handler.HealthPost)
}