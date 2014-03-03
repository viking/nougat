package nougat

import (
	"github.com/neagix/Go-SDL/sdl"
	"testing"
)

var eventTargetTests = []struct {
	kind  uint8
	event Event
}{
	{EvtActive, NewActiveEvent(sdl.ActiveEvent{})},
	{EvtExpose, NewExposeEvent(sdl.ExposeEvent{})},
	{EvtJoyAxis, NewJoyAxisEvent(sdl.JoyAxisEvent{})},
	{EvtJoyBall, NewJoyBallEvent(sdl.JoyBallEvent{})},
	{EvtJoyButton, NewJoyButtonEvent(sdl.JoyButtonEvent{})},
	{EvtJoyHat, NewJoyHatEvent(sdl.JoyHatEvent{})},
	{EvtKeyboard, NewKeyboardEvent(sdl.KeyboardEvent{})},
	{EvtMouseButton, NewMouseButtonEvent(sdl.MouseButtonEvent{})},
	{EvtMouseMotion, NewMouseMotionEvent(sdl.MouseMotionEvent{})},
	{EvtQuit, NewQuitEvent(sdl.QuitEvent{})},
	{EvtResize, NewResizeEvent(sdl.ResizeEvent{})},
	{EvtSysWM, NewSysWMEvent(sdl.SysWMEvent{})},
	{EvtUser, NewUserEvent(sdl.UserEvent{})},
}

func TestDefaultEventTarget(t *testing.T) {
	for i, tt := range eventTargetTests {
		target := new(DefaultEventTarget)
		called := 0
		handler := new(EventListener)
		handler.HandleEvent = func(event Event) {
			called += 1
		}

		target.AddEventListener(tt.kind, handler)
		target.DispatchEvent(tt.event)
		if called != 1 {
			t.Errorf("test %d: handler wasn't called when it should have been", i)
		}

		called = 0
		target.RemoveEventListener(tt.kind, handler)
		target.DispatchEvent(tt.event)
		if called != 0 {
			t.Errorf("test %d: handler was called when it shouldn't have been", i)
		}
	}
}

func TestDefaultEventTarget_Bubbling(t *testing.T) {
	target := new(DefaultEventTarget)
	parent := new(DefaultEventTarget)
	target.eventParent = parent

	parentCalled := 0
	parentHandler := new(EventListener)
	parentHandler.HandleEvent = func(event Event) {
		parentCalled += 1
	}
	parent.AddEventListener(EvtKeyboard, parentHandler)

	event := NewKeyboardEvent(sdl.KeyboardEvent{})
	target.DispatchEvent(event)
	if parentCalled != 1 {
		t.Error("parent handler wasn't called when it should have been")
	}

	parentCalled = 0
	event.StopPropagation()
	target.DispatchEvent(event)
	if parentCalled != 0 {
		t.Error("parent handler was called when it shouldn't have been")
	}
}
