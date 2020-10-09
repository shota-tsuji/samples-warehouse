package main

/* the image package itself doesn't know how to decode jpeg, you need to
import image/jpeg to register the jpeg decoder. (https://github.com/golang/go/issues/9184) */
import (
	//"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	"io"
	"os"
)

func main() {
	src, err := os.Open("src.jpg")
	defer src.Close()
	//img, _ := jpeg.Decode(src)
	//buf := new(bytes.Buffer)
	//jpeg.Encode(buf, img, nil)
	if err != nil {
		panic(err)
	}
	src.Seek(0, io.SeekStart)

	imageData, imageType, err := image.Decode(src)
	//imageData, imageType, err := image.DecodeConfig(src)
	//conf, encode, err := image.DecodeConfig(src)
	//_, _, err = image.DecodeConfig(src)
	if err != nil {
		//panic(err)
	}
	fmt.Println(imageData)
	fmt.Println(imageType)
}
