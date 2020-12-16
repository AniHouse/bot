package app

import "strings"

type Command struct {
	middleware

	aliases []string

	Handler HandlerFunc
}

func (c *Command) Handle(handler HandlerFunc) *Command {
	c.Handler = handler
	return c
}

func (c *Command) match(args *string) bool {
	for _, alias := range c.aliases {
		cut := *args
		for len(cut) != 0 {
			if strings.HasPrefix(cut, alias) {
				cut = strings.TrimPrefix(cut, alias)
				*args = strings.TrimSpace(cut)
				return true
			}
			if cut[0] != ' ' {
				break
			}
			cut = cut[1:]
		}
	}
	return false
}
