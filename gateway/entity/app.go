package entity

type App struct {
	Config Config
}

func NewApp() *App {
	app := &App{Config: NewConfig()}
	return app
}
