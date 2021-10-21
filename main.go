package main

import (
	_ "embed"

	"github.com/TaceyWong/ipic/backend/config"
	"github.com/TaceyWong/ipic/backend/image"
	"github.com/TaceyWong/ipic/backend/stat"

	"github.com/wailsapp/wails"
)

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func main() {
	app := wails.CreateApp(&wails.AppConfig{
		Width:     1200,
		Height:    742,
		Title:     "iPic",
		JS:        js,
		CSS:       css,
		Colour:    "#18181f",
		Resizable: true,
	})

	cfg := config.NewConfig()
	st := stat.NewStat()
	fm := image.NewFileManager(cfg, st)

	app.Bind(cfg)
	app.Bind(st)
	app.Bind(fm)
	_ = app.Run()
}
