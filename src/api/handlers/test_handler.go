package handlers

import ("github.com/gin-gonic/gin"

		"net/http"
)

type TestHandler struct {
}

type header struct{
	UserId string
	Browser string
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"result": "Test",
	})

}

func (h *TestHandler) Users(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"users":"our friends",
	})

}

func (h *TestHandler) UserById(ctx *gin.Context) {

	id:=ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"result": "user",
		"id": id,
	})
}

func (h *TestHandler) Accounts(ctx *gin.Context) {

	id:=ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"result": "accounts",
		"id": id,
	})
}

func (h *TestHandler)AddUser(ctx *gin.Context) {

	id:=ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"result": "add user",
		"id": id,
	})
}

func (h *TestHandler) HeaderBind(ctx *gin.Context) {

	UserId:= ctx.GetHeader("userid")
	ctx.JSON(http.StatusOK, gin.H{
		"result": "HeaderBind",
		"userid": UserId,
	})
}

func (h *TestHandler) HeaderBind2(ctx *gin.Context) {

	header:=header{}
	ctx.JSON(http.StatusOK, gin.H{
		"result": "HeaderBind",
		"header": header,
	})
}