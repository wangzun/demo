package common

var MaxRoleId int
var MaxMapId int

func init() {
	MaxRoleId = 100000000
	MaxMapId = 100
}

func GetNowMapId() int {
	return MaxMapId - 1
}

func InitRoleId(Id int) {
	MaxRoleId = Id + 1
}

func InitMapId(Id int) {
	MaxMapId = Id + 1
}

func GenRoleID() int {
	newId := MaxRoleId
	MaxRoleId++
	return newId
}

func GenMapID() int {
	newId := MaxMapId
	MaxMapId++
	return newId
}
