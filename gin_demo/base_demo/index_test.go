package main

import (
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSomeGetHandler(t *testing.T) {
	router := SetupRoute()
	w := httptest.NewRecorder()

	// 利用  net/http 发请求给我的 gin server
	req, _ := http.NewRequest(http.MethodGet, "/someGet", nil)
	router.ServeHTTP(w, req) // 模拟写的程序处理上面的测试请求，把得到的响应结果写入w

	// 判断得到的响应是否与预期的结果一致
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Hello Gin", w.Body.String())
}
