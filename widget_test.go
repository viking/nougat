package nougat

import (
	"github.com/neagix/Go-SDL/sdl"
)

type fakeWidget struct {
	draw   func() *sdl.Surface
	handle func(event interface{}) bool
}

func (w *fakeWidget) Draw() *sdl.Surface {
	return w.draw()
}

func (w *fakeWidget) Handle(event interface{}) bool {
	return w.handle(event)
}

func (w *fakeWidget) Free() {
}
