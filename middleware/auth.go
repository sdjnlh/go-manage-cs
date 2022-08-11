package middleware

import (
	"github.com/gin-gonic/gin"
	"jinbao-cs/dict"
	"jinbao-cs/model"
	"jinbao-cs/web"
	"net/http"
	"time"
)

// CorsHandler 跨域请求
func CorsHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func JWY(form model.UserInfoDTO, c *gin.Context) {
	var code string

	code = dict.SystemSuccessCode
	if form.Token == "" {
		code = dict.ERROR_AUTH_NO_TOKRN
	} else {
		claims, err := ParseToken(form.Token)
		if err != nil {
			code = dict.ERROR_AUTH_CHECK_TOKEN_FAIL
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = dict.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
		}
	}

	//如果token验证不通过，直接终止程序，c.Abort()
	if code != dict.SystemSuccessCode {
		// 返回错误信息
		result := web.NewFilterResult(&model.User{})
		c.JSON(http.StatusUnauthorized, result)
		//终止程序
		c.Abort()
		return
	}
	c.Next()
}
