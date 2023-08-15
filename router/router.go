package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello gin")
	})
	return router
}
