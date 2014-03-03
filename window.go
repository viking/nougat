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

func NewWindow(w, h int) (win *Window, err error) {
	win = &Window{W: w, H: h}
	win.surface = sdl.SetVideoMode(win.W, win.H, 32, sdl.RESIZABLE)
	if win.surface == nil {
		err = sdlError()
		return
	}

	win.Container = new(Container)
	return
}

func (win *Window) Draw() *sdl.Surface {
	surface := win.Container.Draw()
	win.surface.Blit(nil, surface, nil)
	surface.Free()
	return win.surface
}

func (win *Window) Free() {
	win.surface.Free()
}
