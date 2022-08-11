// Package web
// @Description: 封装请求数据绑定,例如：结构体，int64,int
package web

import (
	"github.com/gin-gonic/gin"
	"jinbao-cs/errors"
	"strconv"
)

type Handler struct {}

func (handler *Handler) Bind(c *gin.Context, domain interface{}) (err error) {
	err = c.ShouldBind(domain)
	if err != nil {
		return error(&errors.SimpleBizError{Code: errors.CommonInvalidParams, Msg: err.Error()})
	}
	return nil
}

func (handler *Handler) Int64Param(c *gin.Context, key string) (int64, error) {
	return strconv.ParseInt(c.Param(key), 10, 64)
}

func (handler *Handler) IntParam(c *gin.Context, key string) (int, error) {
	return strconv.Atoi(c.Param(key))
}

