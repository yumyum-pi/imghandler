package imghandler

import (
	"image"
	// support for reading diffrent image format
	_ "image/jpeg"
	_ "image/png"
	"os"

	// create custom error
	cr "github.com/yumyum-pi/goCreateErr"
)

// Open is a function that opens image from path and return Image and error
func Open(filePath string) (file image.Image, err error) {
	// open file
	reader, err := os.Open(filePath)
	defer reader.Close()
	if err != nil {
		return
	}
	// checking if the file exist
	stat, err := reader.Stat()
	if err != nil {
		return
	}
	if stat.Size() == 0 {
		err = cr.Create("readImg", "The file size is zero", "Use another image")
	}
	// decoding image
	imgData, _, err := image.Decode(reader)
	if err != nil {
		return
	}

	return imgData, err
}
