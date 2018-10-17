package job

import (
	"github.com/robfig/cron"
	"logger"
	"errors"
)

type Scheduler struct {
	Name 	 string
	SpecTime string
	C 		 *cron.Cron
}

func New(name, specTime string) *Scheduler {
	s := new(Scheduler)
	s.Name = name
	s.SpecTime = specTime
	return s
}

//添加定时任务
func (self *Scheduler) AddJob(jobFunc func()) error {
	c := cron.New()
	err := c.AddFunc(self.SpecTime, jobFunc)
	if err != nil {
		return errors.New("启动定时任务:" + self.Name + "失败")
	}
	logger.Info("启动定时任务:%s成功", self.Name)
	self.C = c

	return nil
}

//启动定时任务
func (self *Scheduler) StartJob() {
	self.C.Start()
}

//停止定时任务
func (self *Scheduler) StopJob() {
	self.C.Stop()
}
