package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hosssein21/golang-clean-web-api/api/handlers"
)

func TestRouters(r *gin.RouterGroup){
	h:= handlers.NewTestHandler()

	r.GET("/",h.Test)
	r.GET("/users",h.Users)
	r.GET("/user/:id",h.UserById)
	r.GET("/user/:id/accounts",h.Accounts)
	r.POST("add-user/:id",h.AddUser)
	r.POST("/binder/header1",h.HeaderBind)
	r.POST("/binder/header2",h.HeaderBind2)
}