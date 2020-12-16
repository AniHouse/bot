package test

import (
	"github.com/anihouse/bot/app"
)

func (m *module) onPing(ctx *app.Context) {
	ctx.Session.ChannelMessageSend(
		ctx.Message.ChannelID,
		"pong",
	)
}
