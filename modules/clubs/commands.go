package clubs

import (
	"time"

	"github.com/anihouse/bot/db"

	"github.com/anihouse/bot/app"
	"github.com/anihouse/bot/app/dstype"
	"github.com/anihouse/bot/util"
)

func onClubCreeate(ctx *app.Context) {
	var (
		symbol dstype.Grapheme
		title  string
	)

	err := ctx.Scan(&symbol, &title)
	if err != nil {
		ctx.Error(err)
		return
	}

	expiredAt := util.Midnight(time.Now().UTC().Add(conf.NotVerifiedLifetime))

	club := db.Club{
		OwnerID:   ctx.Message.Author.ID,
		Symbol:    symbol.Value,
		Title:     title,
		ExpiredAt: &expiredAt,
	}

	err = db.Clubs.Create(&club)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.TPL("clubs/created.xml").Data(map[string]interface{}{
		"Prefix":     _module.app.Prefix,
		"Owner":      ctx.Message.Author,
		"Club":       club,
		"MinMembers": conf.MinimumMembers,
		"Price":      int64(conf.Price),
	})
}

func onClubApply(ctx *app.Context) {

}
