package logic

import (
	"errors"
	"fmt"

	"github.com/wangzun/demo/common"
	"github.com/wangzun/gogame/engine/core"
	"github.com/wangzun/gogame/engine/math32"
)

type Role struct {
	Name     string
	Id       int
	Pos      PosData // pos in map
	velocity float32 // linear velocity (m/s)
	rotvel   float32 // rotation velocity (rad/s)
	Model    string  //
	MapId    int
	Node     *core.Node
	commands [CMD_LAST]bool // commands states
	state    int
}

func (r *Role) State() int {
	return r.state
}

func (r *Role) Up() {
	r.commands[CMD_FORWARD] = true
	if r.commands[CMD_RUN] {
		r.state = RUN
	} else {
		r.state = WALK
	}
}

func (r *Role) CancelUp() {
	r.commands[CMD_FORWARD] = false
	r.state = STAY
}

func (r *Role) Down() {
	r.commands[CMD_BACKWARD] = true
}

func (r *Role) CancelDown() {
	r.commands[CMD_BACKWARD] = false
}

func (r *Role) Left() {
	r.commands[CMD_LEFT] = true
}

func (r *Role) CancelLeft() {
	r.commands[CMD_LEFT] = false
}

func (r *Role) Right() {
	r.commands[CMD_RIGHT] = true
}

func (r *Role) CancelRight() {
	r.commands[CMD_RIGHT] = false
}

func (r *Role) Run() {
	r.commands[CMD_RUN] = true
}

func (r *Role) CancelRun() {
	r.commands[CMD_RUN] = false
}

const (
	CMD_FORWARD = iota
	CMD_BACKWARD
	CMD_LEFT
	CMD_RIGHT
	CMD_RUN
	CMD_LAST
)

const (
	STAY = iota
	WALK
	RUN
)

var role_list map[int]*Role

func init() {
	role_list = make(map[int]*Role)
}

func GetRoleList() map[int]*Role {
	return role_list
}

func NewRole() *Role {
	Id := common.GenRoleID()
	name := common.GetFullName()
	role := &Role{Id: Id, Name: name}
	role.velocity = 0.9
	role.rotvel = 0.2
	role.Model = "cowboy"
	role.MapId = common.GetNowMapId()
	pos_list := GetMapEmptyPos(role.MapId)
	index := common.GetRandomNum(len(pos_list))
	pos := pos_list[index]
	role.Pos.X = pos.X
	role.Pos.Y = pos.Y
	role.Node = core.NewNode()
	role.Node.SetPositionVec(GetNorPos(role.Pos.X, role.Pos.Y))
	role.state = STAY
	AddRole(role)
	return role
}

func AddRole(role *Role) error {
	_, ok := role_list[role.Id]
	if ok {
		return nil
	}
	fmt.Println(role.Id)
	role_list[role.Id] = role
	return nil
}

func RoleInfo(Id int) (*Role, error) {
	role, ok := role_list[Id]
	if ok {
		return role, nil
	}
	return nil, errors.New("not found role")
}

func roleLoop() {
	for _, t := range role_list {
		t.Up()
		t.Right()

		if t.commands[CMD_LEFT] || t.commands[CMD_RIGHT] {
			// Calculates angle delta to rotate
			angle := t.rotvel * GetFDS()
			if t.commands[CMD_RIGHT] {
				angle = -angle
			}
			t.Node.RotateY(angle)
		}

		if t.commands[CMD_FORWARD] || t.commands[CMD_BACKWARD] {
			// Calculates the distance to move
			dist := t.velocity * GetFDS()
			// Calculates wheel rotation
			var rot = -dist / 0.5

			// Get tank world direction
			fmt.Println("dist : ", dist)
			var quat math32.Quaternion
			t.Node.WorldQuaternion(&quat)
			// direction := math32.Vector3{1, 0, 0}
			direction := math32.Vector3{0, 0, 1}
			direction.ApplyQuaternion(&quat)
			direction.Normalize()
			direction.MultiplyScalar(dist)
			if t.commands[CMD_BACKWARD] {
				direction.Negate()
				rot = -rot
			}
			// Get tank world position
			var position math32.Vector3
			t.Node.WorldPosition(&position)
			fmt.Println("old position : ", position)
			fmt.Println("dir  : ", direction)
			position.Add(&direction)
			fmt.Println("new position : ", position)
			t.Node.SetPositionVec(&position)
			// Rotate whell caps
		}
	}
}
