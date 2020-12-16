package tpl

import (
	"fmt"
	"html/template"
	"time"

	"github.com/anihouse/bot/util"
	"github.com/bwmarrin/discordgo"
)

var (
	funcs = make(template.FuncMap)
)

func init() {
	funcs["datetime_tz"] = func(t time.Time) string {
		return t.Format("02.01.2006 15:04:05 MST")
	}

	funcs["money"] = func(money int64) string {
		return util.NumberParts(money, " ") + "<:AH_AniCoin:579712087224483850>"
	}

	funcs["mention"] = func(user discordgo.User) string {
		return user.Mention()
	}

	funcs["bold"] = func(s interface{}) string {
		return "**" + fmt.Sprint(s) + "**"
	}
}
