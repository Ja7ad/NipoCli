package module

import (
	"bytes"
	"errors"
	"github.com/Ja7ad/NipoCli/cmd"
	"github.com/Ja7ad/NipoCli/context"
	"github.com/Ja7ad/NipoCli/tools"
	"github.com/Ja7ad/NipoCli/utils"
	"github.com/abiosoft/readline"
	"io"
	"log"
	"sync"
)

// Nipo Promoter
const nipoPromoter = "nipo >>> "
const nipoMultiPromoter = "... "

// Global Errors
var (
	errInput     = errors.New("Incorrect your input, Please try 'help'")
	errInterrupt = errors.New("No interrupt handler")
)

// Shell is interactive struct manage
type Shell struct {
	rootCmd         *cmd.CMD
	generic         func(*context.Context)
	interrupt       func(*context.Context, int, string)
	interruptCount  int
	eof             func(*context.Context)
	reader          *tools.ShellReader
	writer          io.Writer
	active          bool
	activeMutex     sync.RWMutex
	ignoreCase      bool
	customCompleter bool
	haltChan        chan struct{}
	historyFile     string
	autoHelp        bool
	rawArgs         []string
	pager           string
	pagerArgs       []string
	context.ContextValues
	Actions
}

// New Create Nipo Shell with default settings
func New() *Shell {
	return NewConfig(&readline.Config{Prompt: nipoPromoter})
}

func NewConfig(conf *readline.Config) *Shell {
	rl, err := readline.NewEx(conf)
	if err != nil {
		log.Fatal(err)
	}
	return NewReadLine(rl)
}

func NewReadLine(rl *readline.Instance) *Shell {
	shell := &Shell{
		rootCmd: &cmd.CMD{},
		reader: &tools.ShellReader{
			Scanner:     rl,
			Prompt:      rl.Config.Prompt,
			MultiPrompt: nipoMultiPromoter,
			ShowPrompt:  true,
			Buf:         &bytes.Buffer{},
			Completer:   readline.NewPrefixCompleter(),
		},
		writer:   rl.Config.Stdout,
		autoHelp: true,
	}
	shell.Actions = &ShellActions{shell}
	utils.AddDefaultFunc(shell)
	return shell
}
