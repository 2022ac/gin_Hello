package router

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestRouter(t *testing.T) {
	//初始化路由
	router := InitRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	//辅助作用防止报错.
	router.ServeHTTP(w, req)
	//查看状态码是否相同.
	assert.Equal(t, http.StatusOK, w.Code)
	//查看内容是否相同.
	assert.Equal(t, "Hello gin get method", w.Body.String())
}
func TestUserSave(t *testing.T) {
	username := "lisi"
	router := InitRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user/"+username, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户"+username+"已经保存", w.Body.String())
}
func TestUserSaveQuery(t *testing.T) {
	username := "lisi"
	age := 20
	router := InitRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user?name="+username+"&age"+strconv.Itoa(age), nil)
	router.ServeHTTP(w, req)
	//判断状态码
	assert.Equal(t, http.StatusOK, w.Code)
	//判断输出的信息
	assert.Equal(t, "用户"+username+"年龄"+strconv.Itoa(age)+"已经存在", w.Body.String())
}
