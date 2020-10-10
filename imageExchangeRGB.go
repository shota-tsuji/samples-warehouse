package main

import (
	"image"
	//"image/color"
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
			//var rgba [4]uint8
			//var rgba RGBA
			r, g, b, a := m.At(x, y).RGBA()
			r, g, b, a = r>>8, g>>8, b>>8, a>>8
			//r, g, b, a := m.At(x, y).RGBA().(color.RGBA)
			//dstImage.SetRGBA(x, y, color.RGBA{rgba[1], rgba[2], rgba[0], rgba[3]})
			//dstImage.SetRGBA(x, y, rgba)
			stride := dstImage.Stride
			rect := m.Bounds()
			dstImage.Pix[(y-rect.Min.Y)*stride+(x-rect.Min.X)*4] = uint8(g)   // R <- G
			dstImage.Pix[(y-rect.Min.Y)*stride+(x-rect.Min.X)*4+1] = uint8(b) // G <- B
			dstImage.Pix[(y-rect.Min.Y)*stride+(x-rect.Min.X)*4+2] = uint8(r) // B <- R
			dstImage.Pix[(y-rect.Min.Y)*stride+(x-rect.Min.X)*4+3] = uint8(a)
			//dstImage.Pix[(y-rect.Min.Y)*stride+(x-rect.Min.X)*4] = uint8(rgba.G)   // R <- G
			//dstImage.Pix[(y-rect.Min.Y)*stride+(x-rect.Min.X)*4+1] = uint8(rgba.B) // G <- B
			//dstImage.Pix[(y-rect.Min.Y)*stride+(x-rect.Min.X)*4+2] = uint8(rgba.R) // B <- R
			//dstImage.Pix[(y-rect.Min.Y)*stride+(x-rect.Min.X)*4+3] = uint8(rgba.A)
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
