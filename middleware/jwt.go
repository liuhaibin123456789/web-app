package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"web_app/global"
	"web_app/tool"
)

func Jwt() func(c *gin.Context) {

	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")

		if authorization == "" {
			tool.ResponseError(c, global.CodeEmptyAuth)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authorization, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			tool.ResponseError(c, global.CodeInvalidAuth)
			c.Abort()
			return
		}
		// parts[1]是获取到的access_token，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := tool.ParseToken(parts[1])
		if err != nil {
			//如果是过期错误，返回access_token过期错误提示，再次请求时，携带refresh_token，而不是access_token，
			//通过refresh_token拿到新的refresh_token和access_token保存至浏览器，再次请求时，携带这个access_token
			if err == tool.ErrorExpiredToken {
				tool.ResponseError(c, global.CodeExpiredAccessToken)
				c.Abort()
				return
			}
			tool.ResponseError(c, global.CodeWrongToken)
			c.Abort()
			return
		}
		tool.Debug("jwt", zap.Any("mc", mc))

		//校验token是否与redis里的相同
		getToken, err := tool.Get(strconv.FormatInt(mc.UserId, 10))
		if err != nil {
			tool.Debug("tool.RedisGet(strconv.FormatInt(mc.UserId, 10))", zap.Any("error", err))
			tool.ResponseError(c, global.CodeWrongToken)
			c.Abort()
			return
		}
		if getToken != parts[1] {
			tool.Debug("tool.RedisGet(strconv.FormatInt(mc.UserId, 10))", zap.Any("error", errors.New("账号已在另一台设备登录")))
			tool.ResponseError(c, global.CodeUserLoginAlready)
			c.Abort()
			return
		}
		// 将当前请求的user_id、phone信息保存到请求的上下文c上
		c.Set("user_id", mc.UserId)
		c.Next()
	}
}
