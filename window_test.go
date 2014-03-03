package nougat

import (
	"github.com/neagix/Go-SDL/sdl"
	"testing"
)

func TestWindow_Draw(t *testing.T) {
	win, err := NewWindow(20, 20)
	if err != nil {
		t.Fatal(err)
	}

	fake := &fakeWidget{}
	fake.draw = func() *sdl.Surface {
		return sdl.CreateRGBSurface(0, 10, 5, 32, 0, 0, 0, 0)
	}
	win.Add(fake)

	surface := win.Draw()
	if surface.W != 20 {
		t.Errorf("expected surface width to be %d, but was %d", 20, surface.W)
	}
	if surface.H != 20 {
		t.Errorf("expected surface height to be %d, but was %d", 20, surface.H)
	}
}
