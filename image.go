package nougat

import (
	"github.com/neagix/Go-SDL/sdl"
)

type Image struct {
	image *sdl.Surface
}

func NewImage(filename string) (image *Image, err error) {
	var surface *sdl.Surface
	if surface = sdl.Load(filename); surface == nil {
		err = sdlError()
		return
	}

	image = &Image{surface}
	return
}

func NewImageWithColor(color uint32, w int, h int) (image *Image) {
	surface := sdl.CreateRGBSurface(0, w, h, 32, 0, 0, 0, 0)
	surface.FillRect(nil, color)

	image = &Image{surface}
	return
}

func (img *Image) Draw() *sdl.Surface {
	return img.image
}

func (img *Image) Handle(event interface{}) bool {
	// never handle events
	return false
}

func (img *Image) Free() {
	img.image.Free()
}
