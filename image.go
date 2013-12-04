package nougat

import (
	"github.com/neagix/Go-SDL/sdl"
)

type Image struct {
	image *sdl.Surface
}

func NewImage(filename string) (img *Image, err error) {
	var image *sdl.Surface
	if image = sdl.Load(filename); image == nil {
		err = sdlError()
		return
	}

	img = &Image{image}
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
