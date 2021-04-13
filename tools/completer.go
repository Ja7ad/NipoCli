package tools

import (
	"github.com/Ja7ad/NipoCli/cmd"
	"github.com/flynn-archive/go-shlex"
	"strings"
)

type Completer struct {
	cmd      *cmd.CMD
	disabled func() bool
}

func (c Completer) getWords(word []string) (s []string) {
	localCmd, args := c.cmd.FindCommand(word)

	if localCmd == nil {
		localCmd, args = c.cmd, word
	}

	if localCmd.Completer != nil {
		return localCmd.Completer(args)
	}

	for key := range localCmd.SubCommands {
		s = append(s, key)
	}
	return
}

func (c Completer) DoComplete(line []rune, pos int) (newline [][]rune, length int) {
	if c.disabled != nil && c.disabled() {
		return nil, len(line)
	}

	var words []string
	if word, err := shlex.Split(string(line)); err == nil {
		words = word
	} else {
		words = strings.Fields(string(line))
	}

	var compWords []string
	prefix := ""
	if len(words) > 0 && pos > 0 && line[pos-1] != ' ' {
		prefix = words[len(words)-1]
		compWords = c.getWords(words[:len(words)-1])
	} else {
		compWords = c.getWords(words)
	}

	var suggestions [][]rune
	for _, w := range compWords {
		if strings.HasPrefix(w, prefix) {
			suggestions = append(suggestions, []rune(strings.TrimPrefix(w, prefix)))
		}
	}
	if len(suggestions) == 1 && prefix != "" && string(suggestions[0]) == "" {
		suggestions = [][]rune{[]rune(" ")}
	}
	return suggestions, len(prefix)
}
