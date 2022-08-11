package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"jinbao-cs/config"
	"jinbao-cs/cs"
	"jinbao-cs/log"
	"jinbao-cs/middleware"
	"jinbao-cs/newSession"
	"jinbao-cs/rest"
	"jinbao-cs/util"
)

/**
 * @Author: h.li
 * @Date: 2022/03/24
 * @Description: 程序主函数入口，注册后端接口，启动服务
 */
func main() {
	//启动控制台输出格式化
	if err := log.Logger.Starter(); err != nil {
		panic(err)
	}
	cs.SessionMgr = newSession.NewSessionMgr("Cookies", 3600)
	viewServer := gin.New()
	viewServer.Use(gin.Logger())
	viewServer.Use(gin.Recovery())
	viewServer.Use(middleware.CorsHandler())
	api := viewServer.Group("/api")
	//api.Use(security.LogInterceptor)  //记入操作日志表
	rest.RegisterFontAPI(api)
	var g errgroup.Group
	g.Go(func() error {
		return viewServer.Run(":" + cs.ApiConfig.Address)
	})
	if err := g.Wait(); err != nil {
		fmt.Print(err)
	}
}

/**
 * @Author: h.li
 * @Date: 2022/07/01
 * @Description: 在main函数执行前自动被调用，连接数据库，并获取配置文件中程序信息（监听地址、展示标题、版本信息等）
 */
func init() {
	var err error
	//调用数据库工具包，连接数据库
	config := config.Configer()
	cs.Sql, err, cs.ApiConfig = util.BuildDBConnection(config)
	if err != nil {
		log.Logger.Error("fail to load config file message", zap.Error(err))
		return
	}
}
