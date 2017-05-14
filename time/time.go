package main

import (
	"time"

	"github.com/spf13/viper"

	"github.com/pltanton/yags-timer/core"
	"github.com/pltanton/yags/plugins"
)

// New creates timer with a representing time task
func New(conf *viper.Viper) plugins.Plugin {
	conf = setDefaults(conf)
	task := func() string {
		format := conf.GetString("timeFormat")
		return time.Now().Format(format)
	}
	return core.NewTimerFunc(conf, task)
}
