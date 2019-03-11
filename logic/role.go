package logic

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/wangzun/demo/common"
	"github.com/wangzun/gogame/engine/math32"
)

type Role struct {
	Name      string
	Id        int
	Pos       PosData
	Direction *math32.Vector3
	velocity  float32 // linear velocity (m/s)
	rotvel    float32 // rotation velocity (rad/s)
	model     string  //
	MapId     int
}

var role_list map[int]*Role

func init() {
	role_list = make(map[int]*Role)
}

func NewRole() *Role {
	Id := common.GenRoleID()
	name := common.GetFullName()
	role := &Role{Id: Id, Name: name}
	role.velocity = 5.0
	role.rotvel = 0.8
	role.model = "cowboy"
	role.MapId = common.GetNowMapId()
	pos_list := GetMapEmptyPos(role.MapId)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	index := r.Intn(len(pos_list))
	pos := pos_list[index]
	role.Pos.X = pos.X
	role.Pos.Y = pos.Y
	centerPos, err := GetCenterPos(role.MapId)
	v1 := &math32.Vector3{role.Pos.X, role.Pos.Y, 0}
	v2 := &math32.Vector3{centerPos.X, centerPos.Y, 0}
	Dir := v2.Sub(v1).Normalize()
	role.Direction = Dir
	if err != nil {
		panic(err)
	}
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
