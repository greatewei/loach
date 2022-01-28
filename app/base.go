package app

var (
	// Version is describe the service version
	Version bool
	// Help is show all commands help
	Help bool
)

type base struct {
	flag *Flag
}
