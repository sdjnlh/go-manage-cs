package rest

import "github.com/gin-gonic/gin"

func RegisterFontAPI(router gin.IRouter) {
	NewUserAPI().Register(router)
	NewCaptchaAPI().Register(router)
}
