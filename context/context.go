package context

import (
	"github.com/Ja7ad/NipoCli/cmd"
	"github.com/Ja7ad/NipoCli/module"
)

// map for values in the context.
type contextValues map[string]interface{}

type Context struct {
	contextValues
	err     error
	Args    []string
	RowArgs []string
	Cmd     cmd.CMD
	module.Actions
}

// Context error interface
func (c *Context) Err(err error) {
	c.err = err
}

// Context getter
func (c contextValues) Get(key string) interface{} {
	return c[key]
}

// Context setter
func (c *contextValues) Set(key string, value interface{}) {
	if *c == nil {
		*c = make(map[string]interface{})
	}
	(*c)[key] = value
}

// Context Deleter
func (c contextValues) Del(key string) {
	delete(c, key)
}

// return all keys in the context
func (c contextValues) Keys() (keys []string) {
	for key := range c {
		keys = append(keys, key)
	}
	return
}
