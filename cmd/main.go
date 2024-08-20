package main

import (
	"os"
	"thumbnail-yt-loader/loader"
)

func main() {
	resp, err := loader.LoadThumbnail("mbeT8mpmtHA")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("exampl.png", resp, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
