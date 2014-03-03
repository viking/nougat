package nougat

import (
	"github.com/neagix/Go-SDL/sdl"
	"testing"
)

func TestEvent_Bubbles(t *testing.T) {
	event := NewKeyboardEvent(sdl.KeyboardEvent{})
	if !event.Bubbles() {
		t.Error("expected event to bubble")
	}

	event.StopPropagation()
	if event.Bubbles() {
		t.Error("expected event to not bubble")
	}
}
