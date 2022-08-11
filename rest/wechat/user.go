package wechat

import (
	"github.com/gin-gonic/gin"
	"jinbao-cs/cs"
	"jinbao-cs/model"
	"jinbao-cs/service/wechat"
	"strconv"
)

type user int

var User user

func (user) login(c *gin.Context) {
	form := &model.User{}
	if err := c.Bind(form); err != nil {
		c.String(400, "参数错误")
		c.Abort()
	}
	if err := wechat.User.Login(form); err != nil {

		c.String(500, "内部服务器错误")
		c.Abort()
		return
	}
	//创建客户端对应cookie以及在服务器中进行记录
	//var sessionID = cs.SessionMgr.StartSession(c.Writer, c.Request)

	//设置变量值
	//cs.SessionMgr.SetSessionVal(sessionID, "UserInfo", form)
	c.JSON(200, form)
}
func (user) get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	form := &model.User{}
	form.Id = id
	if err := wechat.User.Get(form); err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	c.JSON(200, form)
}
func (user) put(c *gin.Context) {
	form := &model.User{}
	err := c.Bind(form)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	if err := wechat.User.Update(form); err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	c.JSON(200, form)
}
func (user) Register(r *gin.RouterGroup) {
	r.POST("/v1/user/login", User.login)
	r.PUT("/v1/user", cs.SessionMgr.CheckCookieValid, User.put)
	r.GET("/v1/user/:id", cs.SessionMgr.CheckCookieValid, User.get)
}
