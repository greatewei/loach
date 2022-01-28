package app

import (
	"errors"
	"fmt"
	"github.com/greatewei/loach/color"
	"os"
	"strings"
)

// App is global instance
type App struct {
	base
	Name          string
	Version       string
	Logo          string
	Describe      string
	Args          []string
	Cmd           []string
	Commands      map[string]*Command
	commandList   []string
	globalCommand *Command
}

// Init is initialize the application
func Init(fns ...func(app *App)) *App {
	app := &App{
		Commands: map[string]*Command{},
		Args:     []string{},
		base: base{
			flag: NewFlag(),
		},
	}
	app.InitGlobalFlag()
	for _, fn := range fns {
		fn(app)
	}
	return app
}

// InitGlobalFlag is init global flag
func (app *App) InitGlobalFlag() {
	app.globalCommand = &Command{
		base:     app.base,
		Name:     "global variable",
		Describe: "global variable",
		Config: func(c *Command) {
			c.BoolVar(&Version, "version", false, "show the service version, default false", "v")
			c.BoolVar(&Help, "help", false, "display help document, default false", "h")
		},
	}
	// 执行初始化
	app.globalCommand.Config(app.globalCommand)
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
	app.commandList = append(app.commandList, command.Name)
	return true, nil
}

// Run is start service
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
	err := app.flag.parse(params)
	if err != nil {
		fmt.Println(color.Sprint(color.RedText, "Error:"), color.Sprint(color.BlackText, err.Error()))
		return
	}
	app.Args = app.flag.Set.Args()
	app.Cmd = cmd
}

func (app *App) handle() {
	// Check the version
	if Version {
		app.showVersion()
		return
	}
	// View the help documents
	if Help {
		app.showHelp()
		return
	}

	// Execute the command
	app.exec()
	return
}

func (app *App) exec() {
	if len(app.Cmd) == 0 {
		color.Println(color.RedText, "missing command")
		os.Exit(1)
	}
	firCmd := app.Cmd[0]
	if command, ok := app.Commands[firCmd]; ok {
		if command.Fn == nil {
			return
		}
		err := command.Fn(command, app.Args)
		if err != nil {
			color.Println(color.RedText, firCmd+" exec failure")
			os.Exit(1)
		}
	} else {
		color.Println(color.RedText, "command does not exist")
		os.Exit(1)
	}
}

func (app *App) showHelp() {
	app.showVersion()
	app.showUsage()
	app.showGlobal()
	app.showCommands()
}

func (app *App) showVersion() {
	color.Println(color.GreenText, app.Logo)
	color.Println(color.YellowText, "version: "+app.Version)
	color.Println(color.YellowText, "info: "+app.Describe)
	fmt.Println()
}

func (app *App) showUsage() {
	color.Println(color.YellowText, "Usage:")
	color.Println(color.BlackText, " main [global options...] COMMAND [--options ...] [arguments ...]")
	fmt.Println()
}

func (app *App) showGlobal() {
	color.Println(color.YellowText, "Global Options:")
	for _, info := range app.globalCommand.flagSet {
		text := fmt.Sprintf(" -%s, --%-10s %s", info.Alias, info.Name, info.Usage)
		color.Println(color.GreenText, text)
	}
	fmt.Println()
}

func (app *App) showCommands() {
	color.Println(color.YellowText, "Custom Command:")
	for _, commondName := range app.commandList {
		commond := app.Commands[commondName]
		commond.ShowFlagSet()
	}
}
