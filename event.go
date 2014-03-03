package nougat

import (
	"github.com/neagix/Go-SDL/sdl"
)

const (
	_ = iota
	EvtActive
	EvtExpose
	EvtJoyAxis
	EvtJoyBall
	EvtJoyButton
	EvtJoyHat
	EvtKeyboard
	EvtMouseButton
	EvtMouseMotion
	EvtQuit
	EvtResize
	EvtSysWM
	EvtUser
)

type Event interface {
	Target() EventTarget
	StopPropagation()
	Bubbles() bool
}

type CommonEvent struct {
	target          EventTarget
	stopPropagation bool
}

func (e *CommonEvent) Target() EventTarget {
	return e.target
}

func (e *CommonEvent) StopPropagation() {
	e.stopPropagation = true
}

func (e *CommonEvent) Bubbles() bool {
	return !e.stopPropagation
}

type ActiveEvent struct {
	*CommonEvent
	sdl.ActiveEvent
}

func NewActiveEvent(event sdl.ActiveEvent) *ActiveEvent {
	return &ActiveEvent{CommonEvent: new(CommonEvent), ActiveEvent: event}
}

type ExposeEvent struct {
	*CommonEvent
	sdl.ExposeEvent
}

func NewExposeEvent(event sdl.ExposeEvent) *ExposeEvent {
	return &ExposeEvent{CommonEvent: new(CommonEvent), ExposeEvent: event}
}

type JoyAxisEvent struct {
	*CommonEvent
	sdl.JoyAxisEvent
}

func NewJoyAxisEvent(event sdl.JoyAxisEvent) *JoyAxisEvent {
	return &JoyAxisEvent{CommonEvent: new(CommonEvent), JoyAxisEvent: event}
}

type JoyBallEvent struct {
	*CommonEvent
	sdl.JoyBallEvent
}

func NewJoyBallEvent(event sdl.JoyBallEvent) *JoyBallEvent {
	return &JoyBallEvent{CommonEvent: new(CommonEvent), JoyBallEvent: event}
}

type JoyButtonEvent struct {
	*CommonEvent
	sdl.JoyButtonEvent
}

func NewJoyButtonEvent(event sdl.JoyButtonEvent) *JoyButtonEvent {
	return &JoyButtonEvent{CommonEvent: new(CommonEvent), JoyButtonEvent: event}
}

type JoyHatEvent struct {
	*CommonEvent
	sdl.JoyHatEvent
}

func NewJoyHatEvent(event sdl.JoyHatEvent) *JoyHatEvent {
	return &JoyHatEvent{CommonEvent: new(CommonEvent), JoyHatEvent: event}
}

type KeyboardEvent struct {
	*CommonEvent
	sdl.KeyboardEvent
}

func NewKeyboardEvent(event sdl.KeyboardEvent) *KeyboardEvent {
	return &KeyboardEvent{CommonEvent: new(CommonEvent), KeyboardEvent: event}
}

type MouseButtonEvent struct {
	*CommonEvent
	sdl.MouseButtonEvent
}

func NewMouseButtonEvent(event sdl.MouseButtonEvent) *MouseButtonEvent {
	return &MouseButtonEvent{CommonEvent: new(CommonEvent), MouseButtonEvent: event}
}

type MouseMotionEvent struct {
	*CommonEvent
	sdl.MouseMotionEvent
}

func NewMouseMotionEvent(event sdl.MouseMotionEvent) *MouseMotionEvent {
	return &MouseMotionEvent{CommonEvent: new(CommonEvent), MouseMotionEvent: event}
}

type QuitEvent struct {
	*CommonEvent
	sdl.QuitEvent
}

func NewQuitEvent(event sdl.QuitEvent) *QuitEvent {
	return &QuitEvent{CommonEvent: new(CommonEvent), QuitEvent: event}
}

type ResizeEvent struct {
	*CommonEvent
	sdl.ResizeEvent
}

func NewResizeEvent(event sdl.ResizeEvent) *ResizeEvent {
	return &ResizeEvent{CommonEvent: new(CommonEvent), ResizeEvent: event}
}

type SysWMEvent struct {
	*CommonEvent
	sdl.SysWMEvent
}

func NewSysWMEvent(event sdl.SysWMEvent) *SysWMEvent {
	return &SysWMEvent{CommonEvent: new(CommonEvent), SysWMEvent: event}
}

type UserEvent struct {
	*CommonEvent
	sdl.UserEvent
}

func NewUserEvent(event sdl.UserEvent) *UserEvent {
	return &UserEvent{CommonEvent: new(CommonEvent), UserEvent: event}
}
