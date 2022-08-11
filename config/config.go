// Package config
// @Description: 配置文件读取方法包
package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"sync"
)

var Config *viper.Viper

type Pair struct {
	Key    string
	Target interface{}
}

var (
	once sync.Once
)

/**
 * @Author: h.li
 * @Date: 2022/03/24
 * @Description: 配置文件读取方法
 * @param names 配置文件名称
 * @return *viper.Viper
 */
func Configer(names ...string) *viper.Viper {
	once.Do(func() {
		if Config == nil {
			Config = viper.New()
			fn := "cs-api"
			if len(names) > 0 {
				fn = names[0]
			}
			Config.SetConfigName(fn)
			Config.AddConfigPath(".")
			err := Config.ReadInConfig()
			if err != nil {
				log.Fatal(err)
			}

		}
	})
	return Config
}

func LoadConfig(name string, config *viper.Viper) error {
	fmt.Println("控制台日志配置文件： " + name)
	flag.Parse()
	config.SetConfigName(name)
	config.AddConfigPath(".")
	err := config.ReadInConfig()
	if err != nil {
		fmt.Printf("读取控制台日志配置文件出错: %s \n", err)
	}

	return err
}
