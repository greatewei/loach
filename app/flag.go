package app

import (
	"flag"
	"io/ioutil"
)

// Info is describe every command info
type Info struct {
	Name  string
	Alias string
	Usage string
}

// Flag is flag.Flag manage
type Flag struct {
	Set *flag.FlagSet
}

// NewFlag is instantiation object
func NewFlag() *Flag {
	f := &Flag{}
	f.InitFlagSet()
	return f
}

// InitFlagSet is instantiation flag set
func (f *Flag) InitFlagSet() {
	if f.Set != nil {
		return
	}
	f.Set = flag.NewFlagSet("loach", flag.ContinueOnError)
	// disable output internal error message on parse flags
	f.Set.SetOutput(ioutil.Discard)
	// nothing to do ... render usage on after parsed
	f.Set.Usage = func() {}
}

func (f *Flag) parse(args []string) error {
	return f.Set.Parse(args)
}
