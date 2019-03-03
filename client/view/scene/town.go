package scene

import (
	"github.com/wangzun/demo/client/app"
	"github.com/wangzun/gogame/engine/loader/gltf"
)

func init() {
	app.ViewList["town"] = &Town{}
}

type Town struct {
}

func (t *Town) Initialize(app *app.App) {
	// geom := geometry.NewTorus(1, .4, 12, 32, math32.Pi*2)
	// mat := material.NewPhong(math32.NewColor("darkblue"))
	// torusMesh := graphic.NewMesh(geom, mat)
	// torusMesh.SetScale(0.5, 0.5, 0.5)
	// torusMesh.SetPosition(0, 0, 0)
	// // torusMesh.SetRotation(0, 1, 0)
	// app.Scene().Add(torusMesh)

	// gltfjson, err := gltf.ParseJSON("assets/issum_the_town_on_capital_isle/scene.gltf")
	gltfjson, err := gltf.ParseJSON("assets/littlest_tokyo/scene.gltf")
	if err != nil {
		app.Log().Error("gltf json ", err)
	}
	node, err := gltfjson.LoadScene(0)
	if err != nil {
		app.Log().Error("load scene ", err)
	}
	node.GetNode().SetScale(0.1, 0.1, 0.1)
	node.GetNode().SetPositionY(19.5)
	// node.GetNode().SetRotationX(-math.Pi / 2)
	app.Scene().Add(node)
}

func (t *Town) Render(app *app.App) {

}
