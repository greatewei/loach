package app

type Command struct {
	base
	Name     string
	Describe string
	Fn       func(c *Command, args []string) error
	Config   func(c *Command)
}
