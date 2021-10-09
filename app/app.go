package app

import (
	"errors"
	"flag"
	"fmt"
	"github.com/greatewei/loach/color"
	"os"
	"strings"
)

// GlobalFlag is global parameter
type GlobalFlag struct {
	Version bool
	Color   bool
	Help    bool
}

// App is global instance
type App struct {
	base
	Name       string
	Version    string
	Logo       string
	Describe   string
	GlobalFlag *GlobalFlag
	Args       []string
	Cmd        []string
	Commands   map[string]*Command
}

// Init is initialize the application
func Init(fns ...func(app *App)) *App {
	app := &App{
		Commands: map[string]*Command{},
		Args:     []string{},
		base: base{
			Flag: NewFlag(),
		},
		GlobalFlag: &GlobalFlag{},
	}
	app.InitGlobalFlag()
	for _, fn := range fns {
		fn(app)
	}
	return app
}

// InitGlobalFlag is init global flag
func (app *App) InitGlobalFlag() {
	// Version
	app.Flag.BoolVar(&app.GlobalFlag.Version, "v", false, "Display version")
	app.Flag.BoolVar(&app.GlobalFlag.Version, "version", false, "Display version")
	// Color
	app.Flag.BoolVar(&app.GlobalFlag.Color, "c", true, "Display color")
	app.Flag.BoolVar(&app.GlobalFlag.Color, "color", true, "Display color")
	// Help
	app.Flag.BoolVar(&app.GlobalFlag.Help, "h", false, "Display help document")
	app.Flag.BoolVar(&app.GlobalFlag.Help, "help", false, "Display help document")
}

// AddCommand is add Creating Custom Commands and Menus
func (app *App) AddCommand(command *Command) (bool, error) {
	if command.Name == "" {
		return false, errors.New("command is called empty")
	}
	if _, ok := app.Commands[command.Name]; ok {
		return false, errors.New("repeat the command")
	}
	command.base = app.base
	app.Commands[command.Name] = command
	return true, nil
}

func (app *App) Run() {
	// initialize the custom command
	for _, command := range app.Commands {
		if command.Config != nil {
			command.Config(command)
		}
	}
	// parse the input parameters
	app.parse()
	// handle the application
	app.handle()
}

func (app *App) parse() {
	var params []string
	var cmd []string
	isCmd := true
	for _, param := range os.Args[1:] {
		if !strings.Contains(param, "-") && isCmd {
			cmd = append(cmd, param)
		} else {
			isCmd = false
			params = append(params, param)
		}
	}
	err := app.Flag.parse(params)
	if err != nil {
		fmt.Println(color.Sprint(color.RedText, "Error:"), color.Sprint(color.BlackText, err.Error()))
		return
	}
	app.Args = app.Flag.Set.Args()
	app.Cmd = cmd
}

func (app *App) handle() {
	// Check the version
	if app.GlobalFlag.Version {
		app.showVersion()
		return
	}
	// View the help documents
	if app.GlobalFlag.Help {
		app.showHelp()
		return
	}

	// Execute the command
	app.exec()
	return
}

func (app *App) exec() {
	if len(app.Cmd) == 0 {
		panic("missing command")
	}
	firCmd := app.Cmd[0]
	if command, ok := app.Commands[firCmd]; ok {
		err := command.Fn(command, app.Args)
		if err != nil {
			panic(firCmd + " exec failure")
		}
	} else {
		panic("command does not exist")
	}
}

func (app *App) showHelp() {
	app.Flag.Set.VisitAll(func(f *flag.Flag) {
		fmt.Println(f.Name)
	})
}

func (app *App) showVersion() {
	color.Println(color.GreenText, app.Logo)
	color.Println(color.YellowText, "version: "+app.Version)
	color.Println(color.YellowText, "info: "+app.Describe)
}
