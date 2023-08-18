package router

import (
	"Project2/Gin_first/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	eng := gin.Default()
	if mode := gin.Mode(); mode == gin.TestMode {
		eng.LoadHTMLGlob("./../templates/*")
	} else {
		eng.LoadHTMLGlob("templates/*")
	}
	eng.StaticFile("/favicon.ico", "./favicon.ico")
	eng.Static("/statics", "./statics")
	index := eng.Group("/")
	{
		index.Any("", handler.Index)
	}
	//添加user
	userRouter := eng.Group("/user")
	{
		userRouter.GET("/:name", handler.Usersave)
		userRouter.GET("", handler.UserSaveByQuery)
		userRouter.POST("/register", handler.UserRegister)
	}
	return eng
}