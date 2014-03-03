package nougat

import (
	"sync"
)

type EventListener struct {
	HandleEvent func(event Event)
}

type EventTarget interface {
	AddEventListener(kind uint8, handler *EventListener)
	RemoveEventListener(kind uint8, handler *EventListener)
	DispatchEvent(event Event)
	EventParent() EventTarget
}

type DefaultEventTarget struct {
	handlers    map[uint8][]*EventListener
	mutex       sync.Mutex
	eventParent EventTarget
}

func (et *DefaultEventTarget) AddEventListener(kind uint8, handler *EventListener) {
	et.mutex.Lock()
	if et.handlers == nil {
		et.handlers = make(map[uint8][]*EventListener)
	}
	et.handlers[kind] = append(et.handlers[kind], handler)
	et.mutex.Unlock()
}

func (et *DefaultEventTarget) RemoveEventListener(kind uint8, handler *EventListener) {
	et.mutex.Lock()
	if et.handlers[kind] != nil {
		for i, h := range et.handlers[kind] {
			if handler == h {
				handlers := make([]*EventListener, len(et.handlers[kind])-1)
				copy(handlers, et.handlers[kind][0:i])
				copy(handlers, et.handlers[kind][i+1:])
				et.handlers[kind] = handlers
				break
			}
		}
	}
	et.mutex.Unlock()
}

func (et *DefaultEventTarget) DispatchEvent(event Event) {
	et.mutex.Lock()

	var kind uint8
	switch event.(type) {
	case *ActiveEvent:
		kind = EvtActive
	case *ExposeEvent:
		kind = EvtExpose
	case *JoyAxisEvent:
		kind = EvtJoyAxis
	case *JoyBallEvent:
		kind = EvtJoyBall
	case *JoyButtonEvent:
		kind = EvtJoyButton
	case *JoyHatEvent:
		kind = EvtJoyHat
	case *KeyboardEvent:
		kind = EvtKeyboard
	case *MouseButtonEvent:
		kind = EvtMouseButton
	case *MouseMotionEvent:
		kind = EvtMouseMotion
	case *QuitEvent:
		kind = EvtQuit
	case *ResizeEvent:
		kind = EvtResize
	case *SysWMEvent:
		kind = EvtSysWM
	case *UserEvent:
		kind = EvtUser
	}

	if kind > 0 && et.handlers[kind] != nil {
		for _, handler := range et.handlers[kind] {
			handler.HandleEvent(event)
		}
	}
	if event.Bubbles() {
		if parent := et.EventParent(); parent != nil {
			parent.DispatchEvent(event)
		}
	}

	et.mutex.Unlock()
}

func (et *DefaultEventTarget) EventParent() EventTarget {
	return et.eventParent
}
