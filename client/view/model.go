package view

import "github.com/wangzun/demo/client/app"

type IModel interface {
	Initialize(*app.App) // Called once to initialize the demo
	Render(*app.App)     // Called at each frame for animations
}
