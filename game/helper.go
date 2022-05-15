package game

import (
	"github.com/faiface/pixel"
	"image"
	"os"
	"runtime"
)

func GetRelativePath(path string) string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	var separator string
	if runtime.GOOS == "windows" {
		separator = "\\"
	} else {
		separator = "/"
	}

	return cwd + separator + path
}

func LoadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return pixel.PictureDataFromImage(img), nil
}
