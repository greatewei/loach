package main

import (
	"github.com/greatewei/loach/app"
	"github.com/greatewei/loach/progress"
	"time"
)

func main() {
	cli := app.Init(func(app *app.App) {
		app.Name = "loach"
		app.Version = "1.0"
		app.Logo = `
.__                      .__     
|  |   _________    ____ |  |__  
|  |  /  _ \__  \ _/ ___\|  |  \ 
|  |_(  <_> ) __ \\  \___|   Y  \
|____/\____(____  /\___  >___|  /
                \/     \/     \/ `
		app.Describe = "Cli tool"
	})
	type Param struct {
		a string
	}
	param := &Param{}
	_, _ = cli.AddCommand(&app.Command{
		Name:     "test",
		Describe: "这是一个测试方法",
		Fn: func(c *app.Command, args []string) error {
			prog := progress.NewProgress(30, progress.Bar)
			for i := 0; i < 30; i++ {
				time.Sleep(1 * time.Second)
				prog.AddProgress(1)
			}
			return nil
		},
		Config: func(c *app.Command) {
			c.Flag.StringVar(&param.a, "param", "param", "this is param")
		},
	})
	cli.Run()
}
