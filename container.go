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

type Placement struct {
	Widget Widget
	X      int32
	Y      int32
	W      int32
	H      int32
}

func (p *Placement) Contains(x, y int32) bool {
	return x >= p.X && x < (p.X+p.W) && y >= p.Y && y < (p.Y+p.H)
}

func (p *Placement) RelPos(x, y int32) (retx int32, rety int32) {
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
		w, h     int32
	)

	// determine surface dimensions
	c.mutex.Lock()
	for _, child := range c.Children {
		surface := child.Widget.Draw()
		child.W = surface.W
		child.H = surface.H

		surfaces = append(surfaces, surface)
		switch c.Pack {
		case PackLayer:
			child.X = 0
			if surface.W > w {
				w = surface.W
			}
			child.Y = 0
			if surface.H > h {
				h = surface.H
			}
		case PackHorizontal:
			child.X = w
			w += surface.W
			child.Y = 0
			if surface.H > h {
				h = surface.H
			}
		case PackVertical:
			child.X = 0
			if surface.W > w {
				w = surface.W
			}
			child.Y = h
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

func (c *Container) Handle(event interface{}) bool {
	switch e := event.(type) {
	case sdl.MouseButtonEvent:
		// find the widget this pertains to
		for i := len(c.Children) - 1; i >= 0; i-- {
			child := c.Children[i]
			if child.Contains(int32(e.X), int32(e.Y)) {
				// make the event coordinates relative to the widget
				x, y := child.RelPos(int32(e.X), int32(e.Y))
				e.X = uint16(x)
				e.Y = uint16(y)
				if child.Widget.Handle(e) {
					return true
				}
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
