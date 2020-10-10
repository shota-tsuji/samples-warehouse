package main

import (
	"image"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	dstImage := image.NewRGBA(image.Rect(0, 0, 895, 1200))
	dstFile, err := os.Create("dst.jpg")
	defer dstFile.Close()
	if err != nil {
		log.Fatalf("create failed.")
	}

	jpeg.Encode(dstFile, dstImage, nil)
}
