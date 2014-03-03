package nougat

import (
	"github.com/neagix/Go-SDL/sdl"
)

type Widget interface {
	Draw() *sdl.Surface
	Handle(event interface{}) bool
	Free()
}
