package app

import (
	"errors"
	"fmt"
	"github.com/greatewei/loach/color"
	"os"
	"strings"
)

type GlobalFlag struct {
	Version bool
}
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

func Init(fns ...func(app *App)) *App {
	app := &App{
		Commands: map[string]*Command{},
		Args:     []string{},
		base: base{
			Flag: NewFlag(),
		},
		GlobalFlag: &GlobalFlag{},
	}
	// 初始化全局参数
	app.InitGlobalFlag()
	// 初始化app
	for _, fn := range fns {
		fn(app)
	}
	return app
}

func (app *App) InitGlobalFlag() {
	app.Flag.BoolVar(&app.GlobalFlag.Version, "v", false, "Display version")
}

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
	err := app.Flag.Parse(params)
	if err != nil {
		fmt.Println(color.Sprint(color.RedText, "Error:"), color.Sprint(color.BlackText, err.Error()))
		return
	}
	app.Args = app.Flag.Set.Args()
	app.Cmd = cmd
}

func (app *App) handle() {
	// 是否查看版本
	if app.GlobalFlag.Version {
		color.Println(color.GreenText, app.Logo)
		color.Println(color.YellowText, "version: "+app.Version)
		color.Println(color.YellowText, "info: "+app.Describe)
		return
	}

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
	return
}

func (app *App) Run() {
	// 初始化配置
	for _, command := range app.Commands {
		if command.Config != nil {
			command.Config(command)
		}
	}
	// 解析输入参数
	app.parse()
	// 根据命令执行方法
	app.handle()
}
