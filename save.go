package imghandler

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"

	//create custom error
	cr "github.com/yumyum-pi/goCreateErr"
)

func saveJpeg(fileToCreate *os.File, imgData image.Image) error {
	var opt jpeg.Options
	opt.Quality = 50

	err := jpeg.Encode(fileToCreate, imgData, &opt)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func savePng(fileToCreate *os.File, imgData image.Image) error {

	var Enc png.Encoder

	// set the best compression
	Enc.CompressionLevel = -3 //BestCompression

	// ok, write out the data into the new PNG file
	err := Enc.Encode(fileToCreate, imgData)

	return err
}

// Save is a function that will save the file in desired path.
func Save(imgData image.Image, path, fileName, format string) (filePath string, err error) {

	absulutePath := fmt.Sprintf("%v/%v.%v", path, fileName, format)

	// check if directory exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// create directory
		os.MkdirAll(path, 0755)
	}

	// save to is disk
	fileToCreate, err := os.Create(absulutePath)

	if format == "jpeg" || format == "jpg" || format == "JPEG" || format == "JPG" {
		err := saveJpeg(fileToCreate, imgData)
		if err != nil {
			return "", err
		}
	} else if format == "PNG" || format == "png" {
		err := savePng(fileToCreate, imgData)
		if err != nil {
			return "", err
		}
	} else {
		err = cr.Create("Save", "The file format is not valid.", "Use only jpeg/png image format")
	}

	return absulutePath, err
}
