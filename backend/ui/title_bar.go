package ui

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func Minimize(c context.Context) {
	runtime.WindowMinimise(c)
}

func Maximize(c context.Context) {
	runtime.WindowMaximise(c)
}

func Unmaximize(c context.Context) {
	runtime.WindowUnmaximise(c)
}

func Exit(c context.Context) {
	runtime.Quit(c)
}
