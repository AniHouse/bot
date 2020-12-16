package app

import (
	"strings"
	"sync"

	"github.com/bwmarrin/discordgo"
)

type middleware struct {
	middlewares HandlerChain
}

type Module struct {
	middleware

	enabled  bool
	commands []*Command

	Prefix string
	Name   string
}

func (m *Module) On(aliases ...string) *Command {
	cmd := &Command{
		aliases: aliases,
	}
	m.commands = append(m.commands, cmd)
	return cmd
}

func (m *middleware) Use(handlers ...HandlerFunc) {
	m.middlewares.append(handlers...)
}

func (m *Module) Enable() {
	m.enabled = true
}

func (m *Module) Disable() {
	m.enabled = false
}

func (m *Module) findCommands(wg *sync.WaitGroup, s *discordgo.Session, mes *discordgo.Message, ctxs *[]*Context) {
	defer wg.Done()

	if !m.enabled {
		return
	}

	if !strings.HasPrefix(mes.Content, m.Prefix) {
		return
	}

	args := strings.TrimPrefix(mes.Content, m.Prefix)

	match := func(cmd *Command) {
		defer wg.Done()

		if !cmd.match(&args) {
			return
		}

		ctx := &Context{
			module:     m,
			handlers:   m.middlewares,
			params:     make(map[string]interface{}),
			Message:    mes,
			Session:    s,
			ModuleName: m.Name,
		}
		ctx.args.Args = args
		ctx.handlers.append(cmd.middlewares...)
		ctx.handlers.append(cmd.Handler)

		*ctxs = append(*ctxs, ctx)
	}

	wg.Add(len(m.commands))
	for _, cmd := range m.commands {
		go match(cmd)
	}
}
