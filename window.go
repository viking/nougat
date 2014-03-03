package nougat

import (
	"github.com/neagix/Go-SDL/sdl"
)

type Window struct {
	*Container
	W       int
	H       int
	surface *sdl.Surface
}

func NewWindow(c *Container, w, h int) (win *Window, err error) {
	win = &Window{Container: c, W: w, H: h}
	win.surface = sdl.SetVideoMode(win.W, win.H, 32, sdl.RESIZABLE)
	if win.surface == nil {
		err = sdlError()
		return
	}
	return
}

func (win *Window) Draw() {
	surface := win.Container.Draw()
	win.surface.Blit(nil, surface, nil)
	surface.Free()

	win.surface.Flip()
}
