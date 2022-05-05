package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web_app/global"
	"web_app/model"
	"web_app/service"
	"web_app/tool"
)

// Register
// @tags user
// @Summary 注册
// @Accept mpfd
// @Produce json
// @Param phone formData string true "手机号" minlength(11)  maxlength(11)
// @Param password formData string true "密码" minlength(8)  maxlength(16)
// @Success 200 {object} tool.ResJson{data=model.ResRegister}
// @Failure 200 {object} tool.ResJson
// @Router /register [post]
func Register(c *gin.Context) {
	user := new(model.User)
	err := c.ShouldBind(user)
	if err != nil {
		//日志记录
		tool.SugaredWarn("注册api参数有误", err)
		tool.ResponseError(c, global.CodeShouldBindError)
		return
	}
	aToken, rToken, err := service.Register(user)
	if err != nil {
		//日志记录
		tool.SugaredError("service.Register(user)", err)
		tool.ResponseError(c, global.CodeFailedRegister)
		return
	}
	tool.SugaredInfof("注册用户成功. 手机号：%s,user_id:，%s", user.UserId, user.Phone)
	tool.ResponseSuccess(c, model.ResRegister{
		RefreshToken: rToken,
		AccessToken:  aToken,
	})
}

// Login
// @tags user
// @Summary 登录
// @Accept mpfd
// @Produce json
// @Param phone formData string true "手机号" minlength(11)  maxlength(11)
// @Param password formData string true "密码" minlength(8)  maxlength(16)
// @Success 200 {object} tool.ResJson{data=model.ResRegister}
// @Failure 200 {object} tool.ResJson
// @Router /login [post]
func Login(c *gin.Context) {
	user := new(model.User)
	err := c.ShouldBind(user)
	if err != nil {
		//日志记录
		tool.SugaredWarn(
			"登录api参数有误",
			err,
			zap.String("location", "func Login(c *gin.Context)"),
		)
		tool.ResponseError(c, global.CodeShouldBindError)
		return
	}
	aToken, rToken, err := service.Login(user)
	if err != nil {
		//日志记录
		tool.SugaredError("service.Login(user)", err)
		tool.ResponseError(c, global.CodeFailedLogin)
		return
	}
	tool.SugaredInfof("登录用户成功. 手机号：%s,user_id:，%s", user.UserId, user.Phone)
	tool.ResponseSuccess(c, model.ResRegister{
		RefreshToken: rToken,
		AccessToken:  aToken,
	})
}

// GetTokens
// @tags user
// @Summary 获取token
// @Description 该api只在access_token失效时使用，并且请求头携带好refresh_token
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "header auth"
// @Param user_id formData string true "用户user_id"
// @Success 200 {object} tool.ResJson{data=model.ResRegister}
// @Failure 200 {object} tool.ResJson
// @Router /tokens [post]
func GetTokens(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	userId := c.PostForm("user_id")
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		//日志记录
		tool.SugaredWarn("GetTokens参数有误", err)
		tool.ResponseError(c, global.CodeInvalidParam)
		return
	}
	aToken, rToken, err := tool.RefreshToken(id, auth)
	if err != nil {
		//日志记录
		tool.SugaredError("tool.RefreshToken(id, auth)", err)
		tool.ResponseError(c, global.CodeFailedRefreshToken)
		return
	}

	tool.ResponseSuccess(c, model.ResRegister{
		RefreshToken: rToken,
		AccessToken:  aToken,
	})
}
