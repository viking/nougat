package nougat

import (
	"github.com/neagix/Go-SDL/sdl"
	"sync"
)

const (
	PackLayer = iota
	PackHorizontal
	PackVertical
	PackNone
)

type Placement struct {
	Widget Widget
	X      int16
	Y      int16
	W      int16
	H      int16
}

func (p *Placement) Contains(x, y int16) bool {
	return x >= p.X && x < (p.X+p.W) && y >= p.Y && y < (p.Y+p.H)
}

func (p *Placement) RelPos(x, y int16) (retx int16, rety int16) {
	retx = x - p.X
	rety = y - p.Y
	return
}

type Container struct {
	Pack     uint8
	Children []*Placement
	mutex    sync.Mutex
}

func (c *Container) Add(w Widget) {
	c.mutex.Lock()
	c.Children = append(c.Children, &Placement{Widget: w})
	c.mutex.Unlock()
}

func (c *Container) AddWithPosition(w Widget, x, y int16) {
	c.mutex.Lock()
	c.Children = append(c.Children, &Placement{Widget: w, X: x, Y: y})
	c.mutex.Unlock()
}

func (c *Container) Remove(widget Widget) {
	c.mutex.Lock()
	for i, child := range c.Children {
		if widget == child.Widget {
			children := make([]*Placement, len(c.Children)-1)
			copy(children, c.Children[0:i])
			copy(children, c.Children[i+1:])
			c.Children = children
		}
	}
	c.mutex.Unlock()
}

func (c *Container) Draw() (result *sdl.Surface) {
	var (
		surfaces []*sdl.Surface
		w, h     int16
	)

	// determine surface dimensions
	c.mutex.Lock()
	for _, child := range c.Children {
		surface := child.Widget.Draw()
		child.W = int16(surface.W)
		child.H = int16(surface.H)

		surfaces = append(surfaces, surface)
		switch c.Pack {
		case PackLayer:
			child.X = 0
			if child.W > w {
				w = child.W
			}
			child.Y = 0
			if child.H > h {
				h = child.H
			}
		case PackHorizontal:
			child.X = w
			w += child.W
			child.Y = 0
			if child.H > h {
				h = child.H
			}
		case PackVertical:
			child.X = 0
			if child.W > w {
				w = child.W
			}
			child.Y = h
			h += child.H
		case PackNone:
			// X and Y coordinates are fixed
			if (child.X + child.W) > w {
				w = child.X + child.W
			}
			if (child.Y + child.H) > h {
				h = child.Y + child.H
			}
		}
	}
	c.mutex.Unlock()

	// blit surfaces
	if w > 0 && h > 0 {
		result = sdl.CreateRGBSurface(0, int(w), int(h), 32, 0, 0, 0, 0)
		rect := new(sdl.Rect)
		for i, surface := range surfaces {
			placement := c.Children[i]
			rect.X = placement.X
			rect.Y = placement.Y
			result.Blit(rect, surface, nil)
			surface.Free()
		}
	}
	return
}

func (c *Container) Handle(event interface{}) bool {
	var x, y int16

	switch e := event.(type) {
	case sdl.MouseButtonEvent:
		x = int16(e.X)
		y = int16(e.Y)
	case sdl.MouseMotionEvent:
		x = int16(e.X)
		y = int16(e.Y)
	default:
		return false
	}

	// find the widget this pertains to
	for i := len(c.Children) - 1; i >= 0; i-- {
		child := c.Children[i]
		if child.Contains(x, y) {
			// make the event coordinates relative to the widget
			x, y = child.RelPos(x, y)

			var relEvent interface{}
			switch e := event.(type) {
			case sdl.MouseButtonEvent:
				e.X = uint16(x)
				e.Y = uint16(y)
				relEvent = e
			case sdl.MouseMotionEvent:
				e.X = uint16(x)
				e.Y = uint16(y)
				relEvent = e
			}
			if child.Widget.Handle(relEvent) {
				return true
			}
		}
	}

	return false
}

func (c *Container) Free() {
	c.mutex.Lock()
	for _, child := range c.Children {
		child.Widget.Free()
	}
	c.Children = []*Placement(nil)
	c.mutex.Unlock()
}
