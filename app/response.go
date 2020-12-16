package app

type action interface {
	run()
}

type actionsList []action

type contextInstance struct {
	ctx *Context
}

func (al *actionsList) push(a action) {
	*al = append(*al, a)
}

func (c *Context) TPL(tplName string) *actionTPL {
	tpl := &actionTPL{
		tplName: tplName,
		contextInstance: contextInstance{
			ctx: c,
		},
	}

	c.actions.push(tpl)
	return tpl
}
