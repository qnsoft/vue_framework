package ErrorHelper

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/wonderivan/logger"
)

/*
*@检查错误函数
 */
func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
		logger.Warn(err)
		//logger.Trace("this is Trace")
		//logger.Debug("this is Debug")
		//logger.Info("this is Info")
		//logger.Warn("this is Warn")
		//logger.Error("this is Error")
		//logger.Crit("this is Critical")
		//logger.Alert("this is Alert")
		//logger.Emer("this is Emergency")
	}
}

/*
*@信息日志
 */
func LogInfo(args ...interface{}) {
	logger.Info(args)
}

/*
*@错误信息日志
 */
func LogErrInfo(args ...interface{}) {
	logger.Error(args)
}

func init() {
	// 通过配置参数直接配置
	logger.SetLogger(`{"Console": {"level": "DEBG"}}`)
	// 通过配置文件配置
	logger.SetLogger("logs/log.json")
}
