package main

import (
	"image"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	src, err := os.Open("src.jpg")
	defer src.Close()
	if err != nil {
		log.Fatal(err)
	}

	m, _, err := image.Decode(src)
	if err != nil {
		log.Fatal(err)
	}

	dstFile, err := os.Create("dst.jpg")
	defer dstFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	if err = jpeg.Encode(dstFile, m, nil); err != nil {
		log.Fatal(err)
	}
}
