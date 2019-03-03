package main

import (
	"github.com/wangzun/demo/client/app"
	_ "github.com/wangzun/demo/client/view/scene"
)

func main() {
	app := new(app.App)
	app.DirData = "assets"
	app.Start()
}
