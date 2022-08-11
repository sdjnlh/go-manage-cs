// Package log
// @Description: 日志格式化包
package log

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"jinbao-cs/config"
	"sync"
)

/**
 *格式化控制台日志输出
 */
type logger struct {
	*zap.Logger
}

func (l *logger) ErrorE(err error, msg ...interface{}) {
	var s string
	if len(msg) > 0 {
		if fmt.Sprintf("%T", msg[0]) == "string" {
			s = msg[0].(string)
		}
	}
	l.Error(s, zap.Error(err))
}

var (
	Logger = logger{}
	Slog   *zap.SugaredLogger
)

var (
	l    *logger
	once sync.Once
)

const ConfigFileNameLog = "log_config"

func (l *logger) Starter() error {
	var logConfig zap.Config
	var conf *viper.Viper = viper.New()
	var err error

	err = config.LoadConfig(ConfigFileNameLog, conf)
	if err != nil {
		return err
	}
	m := map[string]interface{}{}
	if err = conf.UnmarshalKey("log", &m); err != nil {
		return err
	}

	logcfgs, _ := json.Marshal(m)
	fmt.Println("log config: \n" + string(logcfgs))

	if err := json.Unmarshal(logcfgs, &logConfig); err != nil {
		return err
	}

	if Logger.Logger, err = logConfig.Build(); err != nil {
		return err
	}
	Slog = Logger.Sugar()

	Logger.Info("logger inited")

	return nil
}
