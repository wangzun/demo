package logic

import (
	"errors"

	"github.com/wangzun/demo/common"
	"github.com/wangzun/gogame/engine/math32"
)

const MaxX = 100
const MaxY = 100

type Map struct {
	Id    int
	Poses [MaxX][MaxY]int
}

type PosData struct {
	X float32
	Y float32
}

var maps map[int]*Map

func init() {
	maps = make(map[int]*Map)
}

func NewMap() *Map {
	Id := common.GenMapID()
	map_data := &Map{Id: Id}
	maps[Id] = map_data
	return map_data
}

func GetNorPos(x, y float32) *math32.Vector3 {
	norX := x - MaxX/2
	norY := y - MaxY/2
	return &math32.Vector3{X: norX, Z: norY}
}

func GetCenterPos(Id int) (*PosData, error) {
	_, ok := maps[Id]
	if ok {
		return &PosData{MaxX / 2, MaxY / 2}, nil
	}
	return nil, errors.New("not found map")
}

func MapData(Id int) (*Map, error) {
	map_data, ok := maps[Id]
	if ok {
		return map_data, nil
	}
	return nil, errors.New("not found map")
}

func GetMapEmptyPos(Id int) []*PosData {
	list := make([]*PosData, 0)
	for i := 0; i < MaxX; i++ {
		for j := 0; j < MaxY; j++ {
			list = append(list, &PosData{X: float32(i), Y: float32(j)})
		}
	}
	return list
}
