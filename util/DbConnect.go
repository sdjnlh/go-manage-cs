package util

import (
	"encoding/hex"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"jinbao-cs/dict"
	"jinbao-cs/log"
	"jinbao-cs/model/module"
	"jinbao-cs/password"
	"xorm.io/xorm"
)

/**
 * @Author: h.li
 * @Date: 2022/03/24
 * @Description: 数据库连接方法，加载配置文件中数据库配置项
 * @param config: *viper.Viper 读取到的配置信息
 * @return *xorm.Engine
 * @return error
 * @return *module.CsApi 程序配置信息
 */
func BuildDBConnection(config *viper.Viper) (*xorm.Engine, error, *module.CsApi) {
	conf := &module.ConnectionConfig{}
	api := &module.CsApi{
		Version: "",
		Address: dict.SystemAddress,
		Title:   dict.SystemTitle,
		Token: "",
	}
	err := config.Unmarshal(&conf)
	if err != nil {
		return nil, err, api
	}
	api = &conf.CsApi

	//连接数据库
	masterEngine, err := connect(conf.DbMaster)
	if err != nil {
		log.Logger.Error("fail to connect db of master", zap.Error(err))
		return nil, err, api
	}
	return masterEngine, err, api
}

func connect(db module.DbConfig) (*xorm.Engine, error) {
	decodeString := []byte(dict.SystemAesKey)
	src, err := hex.DecodeString(db.Password)
	if err != nil {
		return nil, err
	}
	pwdByte, err := password.AesDecryptECB(src, decodeString)
	if err != nil {
		return nil, err
	}
	pwd := string(pwdByte)
	//uri: "host=47.100.210.119 user=kpi password=PJsLeaJ_l_qDUyORo0 dbname=kpi220312 sslmode=disable"
	uri := "host=" + db.Host + " user="+db.Username + " password=" + pwd + " dbname=" + db.Dbname + " sslmode=disable"
	engine, err := xorm.NewEngine(db.Type, uri)
	if err != nil {
		return engine, err
	}
	if db.MaxIdle > 0 {
		engine.SetMaxIdleConns(db.MaxIdle)
	}
	if db.MaxOpen > 0 {
		engine.SetMaxOpenConns(db.MaxOpen)
	}
	engine.ShowSQL(db.ShowSql)
	return engine, err
}
