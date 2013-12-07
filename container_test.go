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

func TestContainer_HandleMouseButtonEvent_Layer(t *testing.T) {
	c := &Container{}

	fake_1 := &fakeWidget{}
	fake_1.draw = func() *sdl.Surface {
		return sdl.CreateRGBSurface(0, 10, 5, 32, 0, 0, 0, 0)
	}
	var handled_1 []interface{}
	fake_1.handle = func(event interface{}) bool {
		handled_1 = append(handled_1, event)
		return true
	}
	c.Add(fake_1)

	fake_2 := &fakeWidget{}
	fake_2.draw = func() *sdl.Surface {
		return sdl.CreateRGBSurface(0, 5, 10, 32, 0, 0, 0, 0)
	}
	var handled_2 []interface{}
	fake_2.handle = func(event interface{}) bool {
		handled_2 = append(handled_2, event)
		return true
	}
	c.Add(fake_2)

	surface := c.Draw()
	defer surface.Free()

	event := sdl.MouseButtonEvent{
		Type:   sdl.MOUSEBUTTONUP,
		Button: sdl.BUTTON_LEFT,
		State:  sdl.PRESSED,
		X:      uint16(3),
		Y:      uint16(3),
	}
	if !c.Handle(event) {
		t.Error("expected event to be handled, but wasn't")
	}
	if len(handled_2) != 1 {
		t.Error("widget 2 didn't handle the event")
	}
	if event_1, ok := handled_2[0].(sdl.MouseButtonEvent); ok {
		if event_1.X != uint16(3) {
			t.Errorf("expected event.X to be %d, but was %d", 3, event_1.X)
		}
		if event_1.Y != uint16(3) {
			t.Errorf("expected event.Y to be %d, but was %d", 3, event_1.Y)
		}
	} else {
		t.Errorf("event wasn't the correct type")
	}

	event.X = uint16(8)
	if !c.Handle(event) {
		t.Error("expected event to be handled, but wasn't")
	}
	if len(handled_1) != 1 {
		t.Error("widget 1 didn't handle the event")
	}
	if event_2, ok := handled_1[0].(sdl.MouseButtonEvent); ok {
		if event_2.X != uint16(8) {
			t.Errorf("expected event.X to be %d, but was %d", 3, event_2.X)
		}
		if event_2.Y != uint16(3) {
			t.Errorf("expected event.Y to be %d, but was %d", 3, event_2.Y)
		}
	} else {
		t.Errorf("event wasn't the correct type")
	}
}

func TestContainer_HandleMouseButtonEvent_Horizontal(t *testing.T) {
	c := &Container{Pack: PackHorizontal}

	fake_1 := &fakeWidget{}
	fake_1.draw = func() *sdl.Surface {
		return sdl.CreateRGBSurface(0, 10, 5, 32, 0, 0, 0, 0)
	}
	var handled_1 []interface{}
	fake_1.handle = func(event interface{}) bool {
		handled_1 = append(handled_1, event)
		return true
	}
	c.Add(fake_1)

	fake_2 := &fakeWidget{}
	fake_2.draw = func() *sdl.Surface {
		return sdl.CreateRGBSurface(0, 5, 10, 32, 0, 0, 0, 0)
	}
	var handled_2 []interface{}
	fake_2.handle = func(event interface{}) bool {
		handled_2 = append(handled_2, event)
		return true
	}
	c.Add(fake_2)

	surface := c.Draw()
	defer surface.Free()

	event := sdl.MouseButtonEvent{
		Type:   sdl.MOUSEBUTTONUP,
		Button: sdl.BUTTON_LEFT,
		State:  sdl.PRESSED,
		X:      uint16(3),
		Y:      uint16(3),
	}
	if !c.Handle(event) {
		t.Error("expected event to be handled, but wasn't")
	}
	if len(handled_1) != 1 {
		t.Error("widget 1 didn't handle the event")
	}
	if event_1, ok := handled_1[0].(sdl.MouseButtonEvent); ok {
		if event_1.X != uint16(3) {
			t.Errorf("expected event.X to be %d, but was %d", 3, event_1.X)
		}
		if event_1.Y != uint16(3) {
			t.Errorf("expected event.Y to be %d, but was %d", 3, event_1.Y)
		}
	} else {
		t.Errorf("event wasn't the correct type")
	}

	event.X = uint16(13)
	if !c.Handle(event) {
		t.Error("expected event to be handled, but wasn't")
	}
	if len(handled_2) != 1 {
		t.Error("widget 2 didn't handle the event")
	}
	if event_2, ok := handled_2[0].(sdl.MouseButtonEvent); ok {
		if event_2.X != uint16(3) {
			t.Errorf("expected event.X to be %d, but was %d", 3, event_2.X)
		}
		if event_2.Y != uint16(3) {
			t.Errorf("expected event.Y to be %d, but was %d", 3, event_2.Y)
		}
	} else {
		t.Errorf("event wasn't the correct type")
	}
}

func TestContainer_HandleMouseButtonEvent_Vertical(t *testing.T) {
	c := &Container{Pack: PackVertical}

	fake_1 := &fakeWidget{}
	fake_1.draw = func() *sdl.Surface {
		return sdl.CreateRGBSurface(0, 10, 5, 32, 0, 0, 0, 0)
	}
	var handled_1 []interface{}
	fake_1.handle = func(event interface{}) bool {
		handled_1 = append(handled_1, event)
		return true
	}
	c.Add(fake_1)

	fake_2 := &fakeWidget{}
	fake_2.draw = func() *sdl.Surface {
		return sdl.CreateRGBSurface(0, 5, 10, 32, 0, 0, 0, 0)
	}
	var handled_2 []interface{}
	fake_2.handle = func(event interface{}) bool {
		handled_2 = append(handled_2, event)
		return true
	}
	c.Add(fake_2)

	surface := c.Draw()
	defer surface.Free()

	event := sdl.MouseButtonEvent{
		Type:   sdl.MOUSEBUTTONUP,
		Button: sdl.BUTTON_LEFT,
		State:  sdl.PRESSED,
		X:      uint16(3),
		Y:      uint16(3),
	}
	if !c.Handle(event) {
		t.Error("expected event to be handled, but wasn't")
	}
	if len(handled_1) != 1 {
		t.Error("widget 1 didn't handle the event")
	}
	if event_1, ok := handled_1[0].(sdl.MouseButtonEvent); ok {
		if event_1.X != uint16(3) {
			t.Errorf("expected event.X to be %d, but was %d", 3, event_1.X)
		}
		if event_1.Y != uint16(3) {
			t.Errorf("expected event.Y to be %d, but was %d", 3, event_1.Y)
		}
	} else {
		t.Errorf("event wasn't the correct type")
	}

	event.Y = uint16(8)
	if !c.Handle(event) {
		t.Error("expected event to be handled, but wasn't")
	}
	if len(handled_2) != 1 {
		t.Error("widget 2 didn't handle the event")
	}
	if event_2, ok := handled_2[0].(sdl.MouseButtonEvent); ok {
		if event_2.X != uint16(3) {
			t.Errorf("expected event.X to be %d, but was %d", 3, event_2.X)
		}
		if event_2.Y != uint16(3) {
			t.Errorf("expected event.Y to be %d, but was %d", 3, event_2.Y)
		}
	} else {
		t.Errorf("event wasn't the correct type")
	}
}
