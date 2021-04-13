package cmd

import (
	"bytes"
	"fmt"
	"github.com/Ja7ad/NipoCli/context"
	"github.com/Ja7ad/NipoCli/utils"
	"sort"
	"text/tabwriter"
)

// CMD Command Handler
type CMD struct {
	Name        string
	Aliases     []string
	Func        func(c *context.Context)
	Help        string
	LongHelp    string
	Completer   func(args []string) []string
	subcommands map[string]*CMD
}

type cmdSorter []*CMD

func (c cmdSorter) Len() int           { return len(c) }
func (c cmdSorter) Less(i, j int) bool { return c[i].Name < c[j].Name }
func (c cmdSorter) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

// AddCommand Add New command
func (c *CMD) AddCommand(cmd *CMD) {
	if c.subcommands == nil {
		c.subcommands = make(map[string]*CMD)
	}
	c.subcommands[cmd.Name] = cmd
}

// DeleteCommand Delete command
func (c *CMD) DeleteCommand(name string) {
	delete(c.subcommands, name)
}

// SubCommand return subcommand list
func (c *CMD) SubCommand() []*CMD {
	var cmds []*CMD
	for _, cmd := range c.subcommands {
		cmds = append(cmds, cmd)
	}
	sort.Sort(cmdSorter(cmds))
	return cmds
}

// CheckSubCommand Check your command is sub command
func (c *CMD) CheckSubCommand() bool {
	if len(c.subcommands) > 1 {
		return true
	}
	if _, ok := c.subcommands["help"]; !ok {
		return len(c.subcommands) > 0
	}
	return false
}

// NipoHelp return list commands in sub commands
func (c CMD) NipoHelp() string {
	var b bytes.Buffer
	p := func(s ...interface{}) {
		fmt.Fprintln(&b)
		if len(s) > 0 {
			fmt.Fprintln(&b, s...)
		}
	}
	if c.LongHelp != "" {
		p(c.LongHelp)
	} else if c.Help != "" {
		p(c.Help)
	} else if c.Name != "" {
		p(c.Name, "Nipo Help not available")
	}
	if c.CheckSubCommand() {
		info := fmt.Sprintf("\nNipoCli %s\n%s\nHome: %s\n\nNipo Commands : ", utils.NipoVer, utils.Desc, utils.Web)
		p(info)
		w := tabwriter.NewWriter(&b, 0, 4, 2, ' ', 0)
		for _, child := range c.SubCommand() {
			fmt.Fprintf(w, "\t%s\t\t\t%s\n", child.Name, child.Help)
		}
		w.Flush()
		p()
	}
	return b.String()
}

// FindSubCommand Find sub command in list commands
func (c *CMD) FindSubCommand(name string) *CMD {
	if cmd, ok := c.subcommands[name]; ok {
		return cmd
	}
	for _, cmd := range c.subcommands {
		for _, alias := range cmd.Aliases {
			if alias == name {
				return cmd
			}
		}
	}
	return nil
}

// FindCommand This return commands
func (c CMD) FindCommand(args []string) (*CMD, []string) {
	var cmd *CMD
	for i, arg := range args {
		if command := c.FindSubCommand(arg); command != nil {
			cmd = command
			c = *cmd
			continue
		}
		return cmd, args[i:]
	}
	return cmd, nil
}
