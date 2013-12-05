package nougat

import (
	"github.com/neagix/Go-SDL/sdl"
	"sync"
)

const (
	PackLayer = iota
	PackHorizontal
	PackVertical
)

type Container struct {
	Pack    uint8
	Widgets []Widget
	mutex   sync.Mutex
}

func (c *Container) Draw() (result *sdl.Surface) {
	var (
		surfaces []*sdl.Surface
		w, h     int32
	)

	// determine surface dimensions
	c.mutex.Lock()
	for _, widget := range c.Widgets {
		surface := widget.Draw()
		surfaces = append(surfaces, surface)
		switch c.Pack {
		case PackLayer:
			if surface.W > w {
				w = surface.W
			}
			if surface.H > h {
				h = surface.H
			}
		case PackHorizontal:
			w += surface.W
			if surface.H > h {
				h = surface.H
			}
		case PackVertical:
			if surface.W > w {
				w = surface.W
			}
			h += surface.H
		}
	}
	c.mutex.Unlock()

	// blit surfaces
	if w > 0 && h > 0 {
		result = sdl.CreateRGBSurface(0, int(w), int(h), 32, 0, 0, 0, 0)
		rect := new(sdl.Rect)
		for _, surface := range surfaces {
			result.Blit(rect, surface, nil)
			switch c.Pack {
			case PackHorizontal:
				rect.X += int16(surface.W)
			case PackVertical:
				rect.Y += int16(surface.H)
			}
			surface.Free()
		}
	}
	return
}

func (c *Container) Add(w Widget) {
	c.mutex.Lock()
	c.Widgets = append(c.Widgets, w)
	c.mutex.Unlock()
}

func (c *Container) Remove(w Widget) {
	c.mutex.Lock()
	for i, widget := range c.Widgets {
		if widget == w {
			widgets := make([]Widget, len(c.Widgets)-1)
			copy(widgets, c.Widgets[0:i])
			copy(widgets, c.Widgets[i+1:])
			c.Widgets = widgets
		}
	}
	c.mutex.Unlock()
}

func (c *Container) Free() {
	c.mutex.Lock()
	for _, widget := range c.Widgets {
		widget.Free()
	}
	c.Widgets = []Widget(nil)
	c.mutex.Unlock()
}
