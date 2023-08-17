package router

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

/*func TestRouter(t *testing.T) {
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
}*/
/*func TestUserSave(t *testing.T) {
	username := "lisi"
	router := InitRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user/"+username, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户"+username+"已经保存", w.Body.String())
}*/
/*func TestUserSaveQuery(t *testing.T) {
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
}*/

func TestInitRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)

	username := "tom"
	router := InitRouter()

	t.Run("index", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "hello gin get method", "返回的HTML页面中应该包含 hello gin get method")
	})
	t.Run("UserSave", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/user/"+username, nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, err, nil)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "用户"+username+"已经保存", w.Body.String())
	})
	t.Run("UserSaveQuery", func(t *testing.T) {
		age := "18"
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/user?name="+username+"&age="+age, nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, err, nil)
		assert.Equal(t, w.Code, http.StatusOK)
		assert.Equal(t, "用户:"+username+",年龄:"+age+"已经保存", w.Body.String())
	})
	t.Run("UserSaveWithNotAge", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/user?name="+username, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, err, nil)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "用户:"+username+",年龄:20已经保存", w.Body.String())
	})
}
func TestUserPostForm(t *testing.T) {
	value := url.Values{}
	value.Add("email", "youngxhui@gmail.com")
	value.Add("password", "1234")
	value.Add("passwordAgain", "1234")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/user/register", bytes.NewBufferString(value.Encode()))
}
func 
