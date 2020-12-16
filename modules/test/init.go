package test

import (
	"github.com/anihouse/bot"
)

var _module module

func init() {
	bot.Modules.Register(_module.ID(), &_module)
}
