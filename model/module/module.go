// Package module
// @Description: 该包用于定义与配置文件相关的结构体,程序启动时读取配置文件初始化相应数据
package module

import "time"

type ConnectionConfig struct {
	DbMaster DbConfig
	DbTarget DbConfig
	CsApi
}

type DbConfig struct {
	Type     string `json:"type"`     //数据库类型
	ShowSql  bool   `json:"showSql"`  //是否展示失sql
	Username string `json:"username"` //数据库用户名
	Password string `json:"password"` //密码
	Host     string `json:"host"`     //地址
	Dbname   string `json:"dbname"`   //数据库
	Port     string `json:"port"`     //端口
	MaxIdle  int    `json:"maxIdle"`  //最大连接数
	MaxOpen  int    `json:"maxOpen"`  //最大打开连接数
}

type CsApi struct {
	Version string `json:"version"` //版本信息
	Address string `json:"address"` //程序后台监听地址
	Title   string `json:"title"`   //程序标题
	Token   string `json:"token"`   //token密钥
}

type TCzrz struct {
	Id   int32     `json:",string" form:"id"` //ID
	Czsj time.Time `json:"czsj"`              //操作时间
	Czr  string    `json:"czr"`               //操作人
	Czip string    `json:"czip"`              //操作地址
	Czdx string    `json:"czdx"`              //
	Czlx string    `json:"czlx"`              //操作类型
	Cznr string    `json:"cznr"`              //操作内容
}
