package scene

import (
	"fmt"

	"github.com/wangzun/demo/client/app"
	"github.com/wangzun/gogame/engine/animation"
	"github.com/wangzun/gogame/engine/loader/gltf"
)

func init() {
	app.ViewList["warrior"] = &Warrior{}
}

type Warrior struct {
	anims []*animation.Animation
}

func (t *Warrior) Initialize(app *app.App) {
	gltfjson, err := gltf.ParseJSON("assets/cowboy/scene.gltf")
	if err != nil {
		app.Log().Error("gltf json ", err)
	}
	node, err := gltfjson.LoadScene(0)
	// node, err := gltfjson.LoadNode(3)
	if err != nil {
		app.Log().Error("load scene ", err)
	}
	fmt.Println(node)
	node.GetNode().SetScale(0.001, 0.001, 0.001)
	node.GetNode().SetPosition(0, -0.08, 0)

	app.Scene().Add(node)
	for i := range gltfjson.Animations {
		fmt.Println("iiiiiiiiiiiii : ", i)
		anim, _ := gltfjson.LoadAnimation(i)
		anim.SetLoop(true)
		// anim.SetStart(500)
		t.anims = append(t.anims, anim)
	}

}

func (t *Warrior) Render(app *app.App) {
	t.anims[4].Update(app.FrameDeltaSeconds())
	// for _, anim := range t.anims {
	// 	anim.Update(app.FrameDeltaSeconds())
	// }

}
