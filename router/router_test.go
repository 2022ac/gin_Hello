package router

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
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
	assert.Equal(t, "Hello gin", w.Body.String())
}
