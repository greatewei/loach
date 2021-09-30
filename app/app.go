package app

type App struct {
	Name string
	Version string
	Logo 	string
	Describe string
	GlobalCommand map[string]string
}

type Command struct {
	Name string
	Describe string
	Fn    func(args []string) error
}


func NewApp(fns ...func(app *App)) *App{
	app := &App{
		Name:          "Loach",
		Version:       "1.0",
		Logo:          `.__                      .__     
|  |   _________    ____ |  |__  
|  |  /  _ \__  \ _/ ___\|  |  \ 
|  |_(  <_> ) __ \\  \___|   Y  \
|____/\____(____  /\___  >___|  /
                \/     \/     \/ `,
		Describe:      "Cli tool",
		GlobalCommand: nil,
	}
	for _,fn := range fns {
		fn(app)
	}
	return app
}