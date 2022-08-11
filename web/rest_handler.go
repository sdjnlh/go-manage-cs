package web

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"jinbao-cs/dict"
	"jinbao-cs/errors"
	"jinbao-cs/log"
	"net/http"
)

type RestHandler struct {
	*Handler
}

var DefaultRestHandler = &RestHandler{}

func (handler *RestHandler) Success(c *gin.Context) {
	handler.SuccessWithData(c, nil)
}

func (handler *RestHandler) SuccessWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func (handler *RestHandler) Fail(c *gin.Context) {
	c.AbortWithStatus(http.StatusInternalServerError)
}

func (handler *RestHandler) FailWithCode(c *gin.Context, code string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, errors.SimpleBizError{Code: code})
}

func (handler *RestHandler) FailWithMessage(c *gin.Context, code string, message string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, errors.SimpleBizError{Code: code, Msg: message})
}

func (handler *RestHandler) FailWithError(c *gin.Context, err error) {
	log.Logger.Warn("api fail", zap.Error(err))
	c.AbortWithStatus(http.StatusInternalServerError)
}

func (handler *RestHandler) FailWithBizError(c *gin.Context, err errors.BizError) {
	log.Logger.Warn("api fail", zap.Error(err))
	c.AbortWithStatusJSON(http.StatusInternalServerError, err)
}

func (handler *RestHandler) BadRequest(c *gin.Context) {
	c.AbortWithStatus(http.StatusBadRequest)
}

func (handler *RestHandler) BadRequestWithError(c *gin.Context, err error) {
	log.Logger.Warn("api fail", zap.Error(err))
	c.AbortWithStatusJSON(http.StatusBadRequest, err)
}

func (handler *RestHandler) Unauthorized(c *gin.Context) {
	c.AbortWithStatus(http.StatusUnauthorized)
}

func (handler *RestHandler) Forbidden(c *gin.Context) {
	c.AbortWithStatus(http.StatusForbidden)
}

func (handler *RestHandler) NotFound(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotFound)
}

func (handler *RestHandler) ResultWithError(c *gin.Context, result IResult, err error) {
	if err != nil {
		log.Logger.Warn("api fail", zap.Error(err))
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		handler.Result(c, result)
	}
}

func (handler *RestHandler) Result(c *gin.Context, result IResult) {
	if result == nil {
		c.AbortWithStatus(http.StatusOK)
		return
	}

	if result.IsOk() {
		result.SetMsg(dict.SystemSuccessMsg)
		result.SetCode(dict.SystemSuccessCode)
		result.SetError(nil)
		c.JSON(http.StatusOK, result)
		return
	} else {
		result.SetMsg(dict.SystemFailureMsg)
		result.SetCode(dict.SystemFailureCode)
	}
	if result.Err() != nil {
		switch result.Err().GetCode() {
		case errors.CommonInvalidParams:
			c.AbortWithStatusJSON(http.StatusBadRequest, result.Err())
			break
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, result.Err())
			break
		}

		return
	}

	c.AbortWithStatus(http.StatusInternalServerError)
}
