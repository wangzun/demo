package scene

import (
	"github.com/wangzun/demo/client/app"
	"github.com/wangzun/gogame/engine/loader/gltf"
)

func init() {
	app.ViewList["warrior"] = &Warrior{}
}

type Warrior struct {
}

func (t *Warrior) Initialize(app *app.App) {
	gltfjson, err := gltf.ParseJSON("assets/CesiumMan/glTF/CesiumMan.gltf")
	if err != nil {
		app.Log().Error("gltf json ", err)
	}
	node, err := gltfjson.LoadScene(0)
	// node, err := gltfjson.LoadNode(3)
	if err != nil {
		app.Log().Error("load scene ", err)
	}
	// node.GetNode().SetScale(10, 10, 10)
	// node.GetNode().SetPositionZ(2)

	app.Scene().Add(node)

}

func (t *Warrior) Render(app *app.App) {

}
