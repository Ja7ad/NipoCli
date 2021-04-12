package cmd

// Command Handler
type CMD struct {
	Name    string
	Aliases []string
	Func    func(c *context)
}
