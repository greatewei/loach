package main

import (
	"fmt"
	"time"

	"github.com/greatewei/loach/app"
	"github.com/greatewei/loach/color"
	"github.com/greatewei/loach/interaction"
	"github.com/greatewei/loach/progress"
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
		a int
	}
	param := &Param{}
	_, _ = cli.AddCommand(&app.Command{
		Name:     "prog",
		Describe: "Just a test order",
		Fn: func(c *app.Command, args []string) error {
			// progress
			prog := progress.NewProgress(100, param.a)
			for i := 0; i < 100; i++ {
				time.Sleep(50 * time.Millisecond)
				prog.AddProgress(1)
			}
			return nil
		},
		Config: func(c *app.Command) {
			c.IntVar(&param.a, "type", 0, "progress type", "")
		},
	})
	_, _ = cli.AddCommand(&app.Command{
		Name:     "input",
		Describe: "Input you text",
		Fn: func(c *app.Command, args []string) error {
			// interaction
			ans, _ := interaction.ReadInput("input you name : ")
			fmt.Print("hello ", ans)
			return nil
		},
	})
	_, _ = cli.AddCommand(&app.Command{
		Name:     "color",
		Describe: "Demo preset and custom font colors",
		Fn: func(c *app.Command, args []string) error {
			color.Println(color.GreenText, "preset green")

			orange := color.RGB(255, 128, 0)
			orange.Println("custom RGB orange")

			teal, err := color.Hex("#008080")
			if err != nil {
				return err
			}
			color.PrintlnCustom(teal, "custom hex teal")

			_ = color.Register("brand", color.MustHex("#6C5CE7"))
			if brand, ok := color.Lookup("brand"); ok {
				brand.Println("named custom brand color")
				color.MustHex("#FFEAA7").Background().Println(" custom background ")
			}
			return nil
		},
	})
	cli.Run()
}
