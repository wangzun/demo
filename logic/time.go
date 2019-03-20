package logic

var frameDeltaSeconds float32

func init() {
}

func GetFDS() float32 {
	return frameDeltaSeconds
}

func Loop(fds float32) {
	frameDeltaSeconds = fds
	roleLoop()
}
