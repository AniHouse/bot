package app

import (
	"math"

	"github.com/bwmarrin/discordgo"
)

const abortIndex int8 = math.MaxInt8 / 2

type Context struct {
	args

	module   *Module
	params   map[string]interface{}
	handlers HandlerChain
	index    int8

	actions actionsList
	errors  errorList

	Session    *discordgo.Session
	Message    *discordgo.Message
	ModuleName string
}

func (c *Context) Param(key string) interface{} {
	return c.params[key]
}

func (c *Context) SetParam(key string, value interface{}) {
	c.params[key] = value
}

func (c *Context) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		if c.index == abortIndex {
			return
		}

		handler := c.handlers[c.index]
		handler(c)
		c.index++
	}

	for i := 0; i < len(c.errors); i++ {

	}

	for i := 0; i < len(c.actions); i++ {
		c.actions[i].run()
	}
}

func (c *Context) Abort() {
	c.index = abortIndex
}

func (c *Context) AboutWithError(err error) {
	c.index = abortIndex
	c.errors.push(err)
}

func (c *Context) Error(err error) *Error {
	result := c.errors.push(err)
	return result
}
