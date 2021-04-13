package module

import "github.com/Ja7ad/NipoCli/cmd"

type Actions interface {
	ReadLine() string
	ReadLineErr() (string, error)

	Println(val ...interface{})
	Print(val ...interface{})
	Printf(format string, val ...interface{})
	ShowPagedOut(text string) error
	SetPrompt(promote string)
	SetMultiPrompt(promote string)
	ShowPrompt(show bool)
	Commands() []*cmd.CMD
	HelpInfo() string
	ClearScreen() error
	Stop()
}

type ShellActions struct {
	*Shell
}

// ReadLine read a line from output
func (s *ShellActions) ReadLine() string {

}
