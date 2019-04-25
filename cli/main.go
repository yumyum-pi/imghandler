package main

import (
	"fmt"

	imghandler "../../imghandler"
)

var path = "./resource/img/original/typo.png"
var export = "./resource/img/edited"

func main() {
	file, err := imghandler.Open(path)

	filePath, err := imghandler.Save(file, export, "typo", "png")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(filePath)
}
