package main

import (
	"os/exec"
	"strings"

	"github.com/spf13/viper"

	"github.com/pltanton/yags-timer/core"
	"github.com/pltanton/yags/plugins"
)

// New creates timer plugin with function builded from configured param
func New(conf *viper.Viper) plugins.Plugin {
	return core.NewTimerFunc(
		conf,
		func() string {
			cmd := conf.GetString("cmd")
			res, err := exec.Command(cmd).Output()
			if err != nil {
				return err.Error()
			}
			return strings.TrimSpace(string(res))
		},
	)
}
