package app

import (
	"strings"

	"github.com/anihouse/bot/tpl"
)

type actionTPL struct {
	contextInstance

	tplName string
	data    interface{}

	channelID string
}

func (a *actionTPL) run() {
	send, err := tpl.ToSend(a.tplName, a.data)
	if err != nil {
		return
	}

	var (
		session = a.ctx.Session
	)

	if strings.TrimSpace(a.channelID) == "" {
		a.channelID = a.ctx.Message.ChannelID
	}

	_, err = session.ChannelMessageSendComplex(a.channelID, send)
	if err != nil {
		return
	}
}

func (a *actionTPL) Data(data interface{}) *actionTPL {
	a.data = data
	return a
}

func (a *actionTPL) ChannalID(channelID string) *actionTPL {
	a.channelID = channelID
	return a
}
