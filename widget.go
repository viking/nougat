package nougat

import (
	"errors"
	"github.com/neagix/Go-SDL/sdl"
)

func sdlError() error {
	return errors.New(sdl.GetError())
}

type Widget interface {
	Draw() *sdl.Surface
	Free()
}
