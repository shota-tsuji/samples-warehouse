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

	dstImage := image.NewRGBA(m.Bounds())
	for x := 0; x < m.Bounds().Max.X; x++ {
		for y := 0; y < m.Bounds().Max.Y; y++ {
			//var rgba [4]uint32
			//rgba = m.At(x, y).RGBA()
		}
	}

	dstFile, err := os.Create("dst.jpg")
	defer dstFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	if err = jpeg.Encode(dstFile, dstImage, nil); err != nil {
		log.Fatal(err)
	}
}
