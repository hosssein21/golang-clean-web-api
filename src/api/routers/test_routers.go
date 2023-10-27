package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hosssein21/golang-clean-web-api/api/handlers"
)

func TestRouters(r *gin.RouterGroup){
	h:= handlers.NewTestHandler()

	r.GET("/",h.Test)
}