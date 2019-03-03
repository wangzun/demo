package scene

import (
	"github.com/wangzun/demo/client/app"
	"github.com/wangzun/gogame/engine/loader/gltf"
)

func init() {
	// app.ViewList["bot"] = &Bot{}
}

type Bot struct {
}

func (t *Bot) Initialize(app *app.App) {
	gltfjson, err := gltf.ParseJSON("assets/sphere_bot/scene.gltf")
	if err != nil {
		app.Log().Error("gltf json ", err)
	}
	node, err := gltfjson.LoadNode(3)
	if err != nil {
		app.Log().Error("load scene ", err)
	}
	node.GetNode().SetScale(0.2, 0.2, 0.2)
	node.GetNode().SetPositionZ(-2)

	app.Scene().Add(node)

}

func (t *Bot) Render(app *app.App) {

}
