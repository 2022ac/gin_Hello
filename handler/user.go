package handler

import "github.com/gin-gonic/gin"

func Usersave(context *gin.Context) {
	//请求获取了路径中的参数。
	username := context.Param("name")
	context.String(200, "用户"+username+"已经保存")
}

// 通过query的方法获取参数
func UserSaveByQuery(context *gin.Context) {
	username := context.Query("name")
	age := context.DefaultQuery("age", "20")
	context.String(200, "用户"+username+"年龄"+age+"已经存在")
}
