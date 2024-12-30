package config

import "github.com/uptrace/bun"

type App struct {
	Db *bun.DB
}

func NewApp(db *bun.DB) *App {
	return &App{Db: db}
}
