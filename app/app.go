package app

import (
	"sync"

	"github.com/anihouse/bot/config"

	"github.com/anihouse/bot"

	"github.com/bwmarrin/discordgo"
)

type (
	HandlerFunc  func(*Context)
	HandlerChain []HandlerFunc
)

type Application struct {
	middleware

	modules map[string]*Module
}

func (chain *HandlerChain) append(handlers ...HandlerFunc) {
	*chain = append(*chain, handlers...)
}

func (a *Application) onHandle(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.GuildID != config.Bot.GuildID {
		return
	}

	var ctxs []*Context

	wg := sync.WaitGroup{}
	wg.Add(len(a.modules))
	for _, module := range a.modules {
		module.findCommands(&wg, s, m.Message, &ctxs)
	}
	wg.Wait()

	for _, ctx := range ctxs {
		go func(c *Context) {
			c.index = -1
			c.Next()
		}(ctx)
	}
}

var (
	app *Application = &Application{
		modules: make(map[string]*Module),
	}
)

func Init() {
	bot.Session.AddHandler(app.onHandle)
}

func NewModule(name, prefix string) *Module {
	m := &Module{
		Prefix: prefix,
		Name:   name,
	}
	app.modules[name] = m
	return m
}

func Use(handlers ...HandlerFunc) {
	app.Use(handlers...)
}
