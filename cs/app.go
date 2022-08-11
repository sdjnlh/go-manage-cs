package cs

import (
	"jinbao-cs/model/module"
	"jinbao-cs/newSession"
	"xorm.io/xorm"
)

// 全局变量，用于整个程序,获取监听端口、标题等
var ApiConfig *module.CsApi

var Sql *xorm.Engine
var SessionMgr *newSession.SessionMgr = nil
