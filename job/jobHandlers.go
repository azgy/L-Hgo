package job

import (
	"logger"
	"config/ini"
)

var defalutCron = "0 0 0 * * ?"
var cronTime string

func CreateLogfileByday() {
	ct, err := ini.GetConfig("Log", "logFileByDayTime")
	if err != nil {
		cronTime = defalutCron
		logger.Warn("按天切割日志定时任务执行时间配置有误,设为默认:%s", defalutCron)
	} else {
		cronTime = ct
	}
	scheduler := New("按天切割日志", cronTime)
	err = scheduler.AddJob(func() {
		logger.Info("执行按天切割日志任务开始")
		logger.Init()
		logger.Info("执行按天切割日志任务结束")
	})

	if err != nil {
		logger.Error(err.Error())
	} else {
		scheduler.StartJob()
	}
}