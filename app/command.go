package app

import (
	"fmt"
	"github.com/greatewei/loach/color"
)

// Command is Creating Custom Commands and Menus
type Command struct {
	base
	Name     string
	Describe string
	Fn       func(c *Command, args []string) error
	Config   func(c *Command)
	flagSet  []*Info
}

func (c *Command) setFlagSet(name string, usage string, alias string) {
	if c.flagSet == nil {
		c.flagSet = make([]*Info, 0)
	}
	info := &Info{
		Name:  name,
		Alias: alias,
		Usage: usage,
	}
	c.flagSet = append(c.flagSet, info)
}

// ShowFlagSet is show flag info
func (c *Command) ShowFlagSet() {
	text := c.Name + "  : " + c.Describe + " \n"
	for _, info := range c.flagSet {
		flagText := fmt.Sprintf("%10s-%s, --%-10s %s\n", color.SpacePlaceholder, info.Alias, info.Name, info.Usage)
		text += flagText
	}
	color.Print(color.BlueText, text)
}

// IntVar is set int params
func (c *Command) IntVar(p *int, name string, value int, usage string, alias string) {
	c.flag.Set.IntVar(p, name, value, usage)
	if alias != "" {
		c.flag.Set.IntVar(p, alias, value, usage)
	}
	c.setFlagSet(name, usage, alias)
}

// StringVar is set string params
func (c *Command) StringVar(p *string, name string, value string, usage string, alias string) {
	c.flag.Set.StringVar(p, name, value, usage)
	if alias != "" {
		c.flag.Set.StringVar(p, alias, value, usage)
	}
	c.setFlagSet(name, usage, alias)
}

// BoolVar is set boolean params
func (c *Command) BoolVar(p *bool, name string, value bool, usage string, alias string) {
	c.flag.Set.BoolVar(p, name, value, usage)
	if alias != "" {
		c.flag.Set.BoolVar(p, alias, value, usage)
	}
	c.setFlagSet(name, usage, alias)
}
