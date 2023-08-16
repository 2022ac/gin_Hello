package router

import (
	"Project2/Gin_first/handler"
	"github.com/gin-gonic/gin"
	"strings"
)

// 将同样逻辑的内容抽象成方法
func retHelloGinAndMethod(context *gin.Context) {
	context.String(200, "Hello gin"+" "+strings.ToLower(context.Request.Method)+" method")
}
func InitRouter() *gin.Engine {
	router := gin.Default()
	/*router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin get methed")
	})
	router.POST("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin post method")
	})
	router.PUT("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin put method")
	})
	router.DELETE("/", func(context *gin.Context) {
		context.String(200, "hello gin delete method")
	})
	router.PATCH("/", func(context *gin.Context) {
		context.String(200, "hello gin patch method")
	})
	router.HEAD("/", func(context *gin.Context) {
		context.String(200, "hello gin head method")
	})
	router.OPTIONS("/", func(context *gin.Context) {
		context.String(200, "hello gin OPTIONS method")
	})*/
	/*router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello gin")
	})*/
	//
	index := router.Group("/")
	{
		index.Any("", retHelloGinAndMethod)
	}
	//添加user
	userRouter := router.Group("/user")
	{
		userRouter.GET("/:name", handler.Usersave)
		userRouter.GET("", handler.UserSaveByQuery)
	}
	return router
}
