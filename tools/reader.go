package tools

import (
	"bytes"
	"github.com/abiosoft/readline"
	"strings"
	"sync"
)

type (
	LineString struct {
		line string
		err  error
	}
	ShellReader struct {
		Scanner      *readline.Instance
		Consumers    chan LineString
		Reading      bool
		ReadingMulti bool
		Buf          *bytes.Buffer
		Prompt       string
		MultiPrompt  string
		ShowPrompt   bool
		Completer    readline.AutoCompleter
		sync.Mutex
	}
)

func (s *ShellReader) rlPrompt() string {
	if s.ShowPrompt {
		if s.ReadingMulti {
			return s.MultiPrompt
		}
		return s.Prompt
	}
	return ""
}

func (s *ShellReader) ReadLine(consumer chan LineString) {
	s.Lock()
	defer s.Unlock()

	if s.Reading {
		return
	}
	s.Reading = true
	shellPrompt := s.Prompt
	prompt := s.rlPrompt()
	if s.Buf.Len() > 0 {
		lines := strings.Split(s.Buf.String(), "\n")
		if p := lines[len(lines)-1]; strings.TrimSpace(p) != "" {
			prompt = p
		}
		s.Buf.Truncate(0)
	}
	s.Scanner.SetPrompt(prompt)
	line, err := s.Scanner.Readline()
	s.Scanner.SetPrompt(shellPrompt)
	ls := LineString{string(line), err}
	consumer <- ls
	s.Reading = false
}
