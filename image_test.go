package nougat

import (
	"encoding/base64"
	"github.com/neagix/Go-SDL/sdl"
	"io/ioutil"
	"os"
	"testing"
)

var blueBMP string = "Qk3aAAAAAAAAAIoAAAB8AAAABQAAAAUAAAABABgAAAAAAFAAAAASCwAAEgsAAAAAAAAAAAAAAAD/\n" +
	"AAD/AAD/AAAAAAAA/0JHUnMAAAAAAAAAAFS4HvwAAAAAAAAAAGZmZvwAAAAAAAAAAMT1KP8AAAAA\n" +
	"AAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAA/wAA/wAA/wAA/wAA/wAAAP8AAP8AAP8AAP8AAP8AAAD/\n" +
	"AAD/AAD/AAD/AAD/AAAA/wAA/wAA/wAA/wAA/wAAAP8AAP8AAP8AAP8AAP8AAAA="

func decodeBase64Image(str string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func createTempImage(data string) (path string, err error) {
	var (
		imageData []byte
		tempfile  *os.File
	)

	imageData, err = decodeBase64Image(data)
	if err != nil {
		return
	}

	tempfile, err = ioutil.TempFile("", "nougat")
	if err != nil {
		return
	}

	_, err = tempfile.Write(imageData)
	if err != nil {
		return
	}
	tempfile.Close()
	path = tempfile.Name()

	return
}

func TestImage_Draw(t *testing.T) {
	var (
		image    *Image
		err      error
		filename string
		surface  *sdl.Surface
	)

	filename, err = createTempImage(blueBMP)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(filename)

	image, err = NewImage(filename)
	if err != nil {
		t.Fatal(err)
	}
	defer image.Free()

	surface = image.Draw()
	if surface.W != int32(5) {
		t.Errorf("expected width to be %d, but was %d", 5, surface.W)
	}
	if surface.H != int32(5) {
		t.Errorf("expected height to be %d, but was %d", 5, surface.H)
	}
}

func TestImage_Handle(t *testing.T) {
	var (
		image    *Image
		err      error
		filename string
	)

	filename, err = createTempImage(blueBMP)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(filename)

	image, err = NewImage(filename)
	if err != nil {
		t.Fatal(err)
	}
	defer image.Free()

	var empty struct{}
	if image.Handle(empty) {
		t.Errorf("handle function didn't return false")
	}
}

func TestImage_Draw_SolidColor(t *testing.T) {
	var (
		image   *Image
		surface *sdl.Surface
	)

	image = NewImageWithColor(0x0000ff, 10, 10)
	defer image.Free()

	surface = image.Draw()
	if surface.W != int32(10) {
		t.Errorf("expected width to be %d, but was %d", 10, surface.W)
	}
	if surface.H != int32(10) {
		t.Errorf("expected height to be %d, but was %d", 10, surface.H)
	}
}
