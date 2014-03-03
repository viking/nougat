package nougat

import (
	"errors"
	"github.com/neagix/Go-SDL/sdl"
	"sync"
	"time"
)

func sdlError() error {
	return errors.New(sdl.GetError())
}

type Application struct {
	Window     *Window
	running    bool
	eventQueue []interface{}
	eventMutex sync.Mutex
}

func NewApplication() (app *Application, err error) {
	if sdl.Init(sdl.INIT_EVERYTHING) != 0 {
		err = sdlError()
		return
	}
	app = new(Application)
	return
}

func (app *Application) Run() (err error) {
	app.running = true
	ticker := time.NewTicker(time.Second / 50)
	for app.running {
		select {
		case <-ticker.C:
			app.processEvents()
			app.Window.Draw()
		case event := <-sdl.Events:
			app.QueueEvent(event)
		}
	}
	return
}

func (app *Application) QueueEvent(event interface{}) {
	app.eventMutex.Lock()
	app.eventQueue = append(app.eventQueue, event)
	app.eventMutex.Unlock()
}

func (app *Application) processEvents() {
	app.eventMutex.Lock()

	for _, event := range app.eventQueue {
		switch event.(type) {
		case sdl.QuitEvent:
			app.running = false
		default:
			app.Window.Handle(event)
		}
	}
	app.eventQueue = app.eventQueue[:0]

	app.eventMutex.Unlock()
}
