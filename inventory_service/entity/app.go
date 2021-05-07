package entity

import (
	"gorm.io/gorm"
)

type App struct {
	Config Config
	GormDb *gorm.DB
}

func NewApp() *App {
	app := &App{Config: NewConfig()}
	return app
}
