package cron

import "gopkg.in/robfig/cron.v2"

type Cron struct {
	cronService *cron.Cron
}


func NewCron() *Cron{
	cronService := cron.New()
	return &Cron{cronService: cronService}
}

func (c *Cron) AddFunc(spec string, cmd func()) {
	c.cronService.AddFunc(spec, cmd)
}

func (c *Cron) Start() {
	c.cronService.Start()
}