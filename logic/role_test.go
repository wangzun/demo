package logic

import (
	"testing"
)

func TestRole(t *testing.T) {
	NewMap()
	role := NewRole()
	t.Log(role)
}
