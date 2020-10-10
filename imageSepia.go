package main

import (
	"image"
	//"image/color"
	"image/jpeg"
	"log"
	"os"
)

func sepia(src image.Image, dst *image.RGBA) {
	for x := 0; x < src.Bounds().Max.X; x++ {
		for y := 0; y < src.Bounds().Max.Y; y++ {
			r, g, b, a := src.At(x, y).RGBA()
			r, g, b, a = r>>8, g>>8, b>>8, a>>8
			coef := 0.33
			avg := coef*float64(r) + coef*float64(g) + coef*float64(b)
			stride := dst.Stride
			rect := src.Bounds()
			dst.Pix[(y-rect.Min.Y)*stride+(x-rect.Min.X)*4] = uint8(1.03 * avg)
			dst.Pix[(y-rect.Min.Y)*stride+(x-rect.Min.X)*4+1] = uint8(0.8 * avg)
			dst.Pix[(y-rect.Min.Y)*stride+(x-rect.Min.X)*4+2] = uint8(0.55 * avg)
			dst.Pix[(y-rect.Min.Y)*stride+(x-rect.Min.X)*4+3] = uint8(a)
		}
	}
}

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
	//for x := 0; x < m.Bounds().Max.X; x++ {
	//	for y := 0; y < m.Bounds().Max.Y; y++ {
	//		r, g, b, a := m.At(x, y).RGBA()
	//		r, g, b, a = r>>8, g>>8, b>>8, a>>8
	//		coef := 0.33
	//		avg := coef*float64(r) + coef*float64(g) + coef*float64(b)
	//		stride := dstImage.Stride
	//		rect := m.Bounds()
	//		dstImage.Pix[(y-rect.Min.Y)*stride+(x-rect.Min.X)*4] = uint8(1.03 * avg)
	//		dstImage.Pix[(y-rect.Min.Y)*stride+(x-rect.Min.X)*4+1] = uint8(0.8 * avg)
	//		dstImage.Pix[(y-rect.Min.Y)*stride+(x-rect.Min.X)*4+2] = uint8(0.55 * avg)
	//		dstImage.Pix[(y-rect.Min.Y)*stride+(x-rect.Min.X)*4+3] = uint8(a)
	//	}
	//}
	sepia(m, dstImage)

	dstFile, err := os.Create("dst.jpg")
	defer dstFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	if err = jpeg.Encode(dstFile, dstImage, nil); err != nil {
		log.Fatal(err)
	}
}
