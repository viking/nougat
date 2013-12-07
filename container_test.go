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
	c := new(Container)

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

var eventTests = []struct {
	pack          uint8
	event         interface{}
	handledBy     uint
	expectedEvent interface{}
}{
	{PackLayer, sdl.MouseButtonEvent{X: 8, Y: 3}, 0, sdl.MouseButtonEvent{X: 8, Y: 3}},
	{PackLayer, sdl.MouseButtonEvent{X: 3, Y: 3}, 1, sdl.MouseButtonEvent{X: 3, Y: 3}},
	{PackHorizontal, sdl.MouseButtonEvent{X: 3, Y: 3}, 0, sdl.MouseButtonEvent{X: 3, Y: 3}},
	{PackHorizontal, sdl.MouseButtonEvent{X: 13, Y: 3}, 1, sdl.MouseButtonEvent{X: 3, Y: 3}},
	{PackVertical, sdl.MouseButtonEvent{X: 3, Y: 3}, 0, sdl.MouseButtonEvent{X: 3, Y: 3}},
	{PackVertical, sdl.MouseButtonEvent{X: 3, Y: 8}, 1, sdl.MouseButtonEvent{X: 3, Y: 3}},
}

func TestContainer_Handle(t *testing.T) {
	for testNum, tt := range eventTests {
		// setup
		c := &Container{Pack: tt.pack}

		var handled [2][]interface{}
		child_1 := &fakeWidget{}
		child_1.draw = func() *sdl.Surface {
			return sdl.CreateRGBSurface(0, 10, 5, 32, 0, 0, 0, 0)
		}
		child_1.handle = func(event interface{}) bool {
			handled[0] = append(handled[0], event)
			return true
		}
		c.Add(child_1)

		child_2 := &fakeWidget{}
		child_2.draw = func() *sdl.Surface {
			return sdl.CreateRGBSurface(0, 5, 10, 32, 0, 0, 0, 0)
		}
		child_2.handle = func(event interface{}) bool {
			handled[1] = append(handled[1], event)
			return true
		}
		c.Add(child_2)

		surface := c.Draw()
		defer surface.Free()

		// send event
		if !c.Handle(tt.event) {
			t.Errorf("test %d: expected event to be handled, but wasn't", testNum)
		}
		if len(handled[tt.handledBy]) != 1 {
			t.Errorf("test %d: widget %d didn't handle the event", testNum, tt.handledBy)
		} else if event := handled[tt.handledBy][0]; event != tt.expectedEvent {
			t.Errorf("test %d: expected %+v, got %+v", testNum, tt.expectedEvent, event)
		}
	}
}
