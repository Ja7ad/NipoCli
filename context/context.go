package context

import (
	"github.com/Ja7ad/NipoCli/cmd"
	"github.com/Ja7ad/NipoCli/module"
)

// map for values in the context.
type ContextValues map[string]interface{}

type Context struct {
	ContextValues
	err     error
	Args    []string
	RowArgs []string
	Cmd     cmd.CMD
	module.Actions
}

// Err Context error interface
func (c *Context) Err(err error) {
	c.err = err
}

// Get Context getter
func (c ContextValues) Get(key string) interface{} {
	return c[key]
}

// Set Context setter
func (c *ContextValues) Set(key string, value interface{}) {
	if *c == nil {
		*c = make(map[string]interface{})
	}
	(*c)[key] = value
}

// Del Context Deleter
func (c ContextValues) Del(key string) {
	delete(c, key)
}

// Keys return all keys in the context
func (c ContextValues) Keys() (keys []string) {
	for key := range c {
		keys = append(keys, key)
	}
	return
}
