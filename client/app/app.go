package app

import (
	"github.com/wangzun/gogame/engine/gui"
	"github.com/wangzun/gogame/engine/util/application"
)

type App struct {
	*application.Application            // Embedded standard application object
	DirData                  string     // full path of data directory
	labelFPS                 *gui.Label // header FPS label
}
