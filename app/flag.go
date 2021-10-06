package app

import (
	"flag"
	"io/ioutil"
)

type Flag struct {
	Set *flag.FlagSet
}

func NewFlag() *Flag {
	flag := &Flag{}
	flag.InitFlagSet()
	return flag
}

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

func (f *Flag) Parse(args []string) error {
	err := f.Set.Parse(args)
	return err
}

func (f *Flag) IntVar(p *int, name string, value int, usage string) {
	f.Set.IntVar(p, name, value, usage)
}

func (f *Flag) StringVar(p *string, name string, value string, usage string) {
	f.Set.StringVar(p, name, value, usage)
}

func (f *Flag) BoolVar(p *bool, name string, value bool, usage string) {
	f.Set.BoolVar(p, name, value, usage)
}
