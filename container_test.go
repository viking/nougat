package nougat

import (
	"github.com/neagix/Go-SDL/sdl"
	"testing"
)

func TestContainer_Draw_NoWidgets(t *testing.T) {
	c := &Container{}
	surface := c.Draw()
	if surface != nil {
		surface.Free()
		t.Error("expected draw function to return nil")
	}
}

func TestContainer_Draw_OneWidget(t *testing.T) {
	c := &Container{}

	fake := &fakeWidget{}
	fake.draw = func() *sdl.Surface {
		return sdl.CreateRGBSurface(0, 10, 5, 32, 0, 0, 0, 0)
	}
	c.Add(fake)

	surface := c.Draw()
	defer surface.Free()
	if surface.W != 10 {
		t.Errorf("expected surface width to be %d, but was %d", 10, surface.W)
	}
	if surface.H != 5 {
		t.Errorf("expected surface height to be %d, but was %d", 5, surface.H)
	}
}

func TestContainer_Draw_TwoWidgets(t *testing.T) {
	c := &Container{}

	fake_1 := &fakeWidget{}
	fake_1.draw = func() *sdl.Surface {
		return sdl.CreateRGBSurface(0, 10, 5, 32, 0, 0, 0, 0)
	}
	c.Add(fake_1)

	fake_2 := &fakeWidget{}
	fake_2.draw = func() *sdl.Surface {
		return sdl.CreateRGBSurface(0, 5, 10, 32, 0, 0, 0, 0)
	}
	c.Add(fake_2)

	surface := c.Draw()
	defer surface.Free()
	if surface.W != 10 {
		t.Errorf("expected surface width to be %d, but was %d", 10, surface.W)
	}
	if surface.H != 10 {
		t.Errorf("expected surface height to be %d, but was %d", 10, surface.H)
	}
}

func TestContainer_PackHorizontal(t *testing.T) {
	c := &Container{Pack: PackHorizontal}

	fake_1 := &fakeWidget{}
	fake_1.draw = func() *sdl.Surface {
		return sdl.CreateRGBSurface(0, 10, 5, 32, 0, 0, 0, 0)
	}
	c.Add(fake_1)

	fake_2 := &fakeWidget{}
	fake_2.draw = func() *sdl.Surface {
		return sdl.CreateRGBSurface(0, 5, 10, 32, 0, 0, 0, 0)
	}
	c.Add(fake_2)

	surface := c.Draw()
	defer surface.Free()
	if surface.W != 15 {
		t.Errorf("expected surface width to be %d, but was %d", 15, surface.W)
	}
	if surface.H != 10 {
		t.Errorf("expected surface height to be %d, but was %d", 10, surface.H)
	}
}

func TestContainer_PackVertical(t *testing.T) {
	c := &Container{Pack: PackVertical}

	fake_1 := &fakeWidget{}
	fake_1.draw = func() *sdl.Surface {
		return sdl.CreateRGBSurface(0, 10, 5, 32, 0, 0, 0, 0)
	}
	c.Add(fake_1)

	fake_2 := &fakeWidget{}
	fake_2.draw = func() *sdl.Surface {
		return sdl.CreateRGBSurface(0, 5, 10, 32, 0, 0, 0, 0)
	}
	c.Add(fake_2)

	surface := c.Draw()
	defer surface.Free()
	if surface.W != 10 {
		t.Errorf("expected surface width to be %d, but was %d", 10, surface.W)
	}
	if surface.H != 15 {
		t.Errorf("expected surface height to be %d, but was %d", 15, surface.H)
	}
}

func TestContainer_RemoveWidget(t *testing.T) {
	c := &Container{}

	fake_1 := &fakeWidget{}
	fake_1.draw = func() *sdl.Surface {
		return sdl.CreateRGBSurface(0, 10, 5, 32, 0, 0, 0, 0)
	}
	c.Add(fake_1)

	fake_2 := &fakeWidget{}
	fake_2.draw = func() *sdl.Surface {
		return sdl.CreateRGBSurface(0, 5, 10, 32, 0, 0, 0, 0)
	}
	c.Add(fake_2)

	c.Remove(fake_1)

	surface := c.Draw()
	defer surface.Free()
	if surface.W != 5 {
		t.Errorf("expected surface width to be %d, but was %d", 5, surface.W)
	}
	if surface.H != 10 {
		t.Errorf("expected surface height to be %d, but was %d", 10, surface.H)
	}
}

func TestContainer_Handle(t *testing.T) {
	c := new(Container)
	defer c.Free()

	var empty struct{}
	if c.Handle(empty) {
		t.Errorf("handle function didn't return false")
	}
}
