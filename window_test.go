package nougat

import (
	"github.com/neagix/Go-SDL/sdl"
	"testing"
)

func TestWindow(t *testing.T) {
	c := &Container{}

	fake := &fakeWidget{}
	fake.draw = func() *sdl.Surface {
		return sdl.CreateRGBSurface(0, 10, 5, 32, 0, 0, 0, 0)
	}
	c.Add(fake)

	win, err := NewWindow(c, 20, 20)
	if err != nil {
		t.Fatal(err)
	}
	defer win.Free()
	win.Draw()
}
