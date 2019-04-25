package main

import (
	"fmt"
	"image"
	"image/jpeg"
	jpg "image/jpeg"
	_ "image/png"
	"os"

	cr "github.com/yumyum-pi/goCreateErr"
)

const filePath string = "./main.go"

var path string = "./resource/img/original/typo.png"

func readImg(filePath string) (file image.Image, format string, err error) {
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
		err = cr.Create("main.go/readImg", "The file size is zero", "Use another image")
	}
	// decoding image
	imgData, imgType, err := image.Decode(reader)
	if err != nil {
		return
	}

	return imgData, imgType, err
}

func moveImg() bool {
	// read from the file
	imgData, format, err := readImg(path)
	fmt.Println(format)
	if err != nil {
		fmt.Println(err)
		return false
	}
	// crop image
	// save to is disk
	fileToCreate, err := os.Create("./resource/img/edited/typo.jpeg")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return false
	}
	var opt jpeg.Options
	opt.Quality = 50

	err = jpg.Encode(fileToCreate, imgData, &opt)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true

}

func main() {
	fmt.Println(moveImg())
}
