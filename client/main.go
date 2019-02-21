package main

import (
	"github.com/wangzun/demo/client/app"
	"github.com/wangzun/gogame/engine/util/application"
)

func main() {
	app := new(app.App)
	a, _ := application.Create(application.Options{})
	app.Application = a
	app.DirData = "assets"
	app.Run()
}
