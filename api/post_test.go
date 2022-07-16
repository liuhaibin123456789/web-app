package api_test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"web_app/api"
	"web_app/global"
	"web_app/settings"
)

func TestCreatePost(t *testing.T) {
	//初始化依赖项
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/v1/post"
	settings.Settings()
	r.POST(url, api.CreatePost)

	body := `{
    "title":"test",
    "content":"testing the api",
	"community_id":"123"
	}`
	req := httptest.NewRequest(http.MethodPost, url, strings.NewReader(body))
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code, "应该相等")
	assert.Equal(t, w.Body.String(), global.CodeInvalidParam, "应该相等")
}
