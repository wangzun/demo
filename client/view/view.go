package view

var ViewList = map[string]IModel{}

type IModel interface {
	Initialize() // Called once to initialize the demo
	Render()     // Called at each frame for animations
}
