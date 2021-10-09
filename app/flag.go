package app

import (
	"flag"
	"io/ioutil"
)

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

// IntVar is set int flag
func (f *Flag) IntVar(p *int, name string, value int, usage string) {
	f.Set.IntVar(p, name, value, usage)
}

// StringVar is set string flag
func (f *Flag) StringVar(p *string, name string, value string, usage string) {
	f.Set.StringVar(p, name, value, usage)
}

// BoolVar is boolean flag
func (f *Flag) BoolVar(p *bool, name string, value bool, usage string) {
	f.Set.BoolVar(p, name, value, usage)
}

func (f *Flag) parse(args []string) error {
	err := f.Set.Parse(args)
	return err
}
