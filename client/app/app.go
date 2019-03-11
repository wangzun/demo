package app

import (
	"github.com/wangzun/gogame/engine/graphic"
	"github.com/wangzun/gogame/engine/gui"
	"github.com/wangzun/gogame/engine/light"
	"github.com/wangzun/gogame/engine/math32"
	"github.com/wangzun/gogame/engine/util/application"
)

var ViewList = map[string]IModel{}

type IModel interface {
	Initialize(*App) // Called once to initialize the demo
	Render(*App)     // Called at each frame for animations
}

type App struct {
	*application.Application            // Embedded standard application object
	DirData                  string     // full path of data directory
	labelFPS                 *gui.Label // header FPS label
}

func (app *App) Start() {
	a, _ := application.Create(application.Options{Control: true})
	app.Application = a
	ambientLight := light.NewAmbient(&math32.Color{1.0, 1.0, 1.0}, 0.8)
	app.Scene().Add(ambientLight)
	pointLight := light.NewPoint(&math32.Color{1, 1, 1}, 5.0)
	pointLight.SetPosition(10, 10, 10)
	app.Scene().Add(pointLight)
	grid := graphic.NewGridHelper(100, 1, &math32.Color{0.4, 0.4, 0.4})
	a.Scene().Add(grid)
	// // fmt.Println(ViewList)
	app.InitView()
	app.Subscribe(application.OnBeforeRender, func(evname string, ev interface{}) {
		app.Render()
	})
	app.CameraPersp().SetPosition(0, 0, 3)
	app.CameraPersp().LookAt(&math32.Vector3{0, 0, 0})
	app.Run()
}

func (app *App) InitView() {
	for _, v := range ViewList {
		v.Initialize(app)
	}
}

func (app *App) Render() {
	for _, v := range ViewList {
		v.Render(app)
	}
}
