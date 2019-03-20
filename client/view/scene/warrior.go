package scene

import (
	"fmt"

	"github.com/wangzun/demo/client/app"
	"github.com/wangzun/demo/logic"
	"github.com/wangzun/gogame/engine/animation"
	"github.com/wangzun/gogame/engine/core"
	"github.com/wangzun/gogame/engine/loader/gltf"
)

func init() {
	role := make(map[int]*RoleInfo)
	app.ViewList["warrior"] = &Warrior{role: role}
}

type Warrior struct {
	role map[int]*RoleInfo
}

type RoleInfo struct {
	Id    int
	anims []*animation.Animation
	node  *core.Node
}

func (t *Warrior) Initialize(app *app.App) {
	roleList := logic.GetRoleList()
	fmt.Println(roleList)
	for _, v := range roleList {
		t.role[v.Id] = &RoleInfo{Id: v.Id}
		t.role[v.Id].anims = make([]*animation.Animation, 0)
		gltfjson, err := gltf.ParseJSON("assets/" + v.Model + "/scene.gltf")
		if err != nil {
			app.Log().Error("gltf json ", err)
		}
		node, err := gltfjson.LoadScene(0)
		if err != nil {
			app.Log().Error("load scene ", err)
		}
		t.role[v.Id].node = node.GetNode()
		position := v.Node.Position()
		node.GetNode().SetPositionVec(&position)
		rotation := v.Node.Rotation()
		node.GetNode().SetRotationVec(&rotation)

		// for view
		if v.Model == "cowboy" {
			node.GetNode().SetScale(0.001, 0.001, 0.001)
			node.GetNode().SetPositionY(-0.08)
		}
		//
		app.Scene().Add(node)
		for i := range gltfjson.Animations {
			// fmt.Println("iiiiiiiiiiiii : ", i)
			anim, _ := gltfjson.LoadAnimation(i)
			anim.SetLoop(true)
			// anim.SetStart(500)
			t.role[v.Id].anims = append(t.role[v.Id].anims, anim)
		}
	}
}

func (t *Warrior) Render(app *app.App) {
	roleList := logic.GetRoleList()
	for _, v := range t.role {
		roleInfo := roleList[v.Id]
		position := roleInfo.Node.Position()
		v.node.GetNode().SetPositionVec(&position)
		rotation := roleInfo.Node.Rotation()
		v.node.GetNode().SetRotationVec(&rotation)

		state := roleInfo.State()
		if state == logic.STAY {
		} else if state == logic.WALK {
			v.anims[3].Update(app.FrameDeltaSeconds())
		} else if state == logic.RUN {
			v.anims[2].Update(app.FrameDeltaSeconds())
		}
	}
	// t.anims[4].Update(app.FrameDeltaSeconds())
	// for _, anim := range t.anims {
	// 	anim.Update(app.FrameDeltaSeconds())
	// }
}
