package rest

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"jinbao-cs/log"
	"jinbao-cs/middleware"
	"jinbao-cs/model"
	"jinbao-cs/web"
)

type CaptchaAPI struct {
	*web.RestHandler
}

func NewCaptchaAPI() *CaptchaAPI {
	return &CaptchaAPI{
		web.DefaultRestHandler,
	}
}

func (api *CaptchaAPI) GetCaptcha(c *gin.Context) {
	id, base64s, err := middleware.GenerateCaptcha()
	if err != nil {
		log.Logger.Error("fail to get captcha", zap.Error(err))
	}
	result := web.NewFilterResult(&model.Captcha{})
	captcha := &model.Captcha{
		CaptchaId: id,
		CaptchaImg:   base64s,
	}
	result.Data = captcha
	result.Ok = true
	api.ResultWithError(c, result, err)
}
func (api *CaptchaAPI) Register(router gin.IRouter) {
	v1 := router.Group("/captcha")
	v1.GET("", api.GetCaptcha)
}
