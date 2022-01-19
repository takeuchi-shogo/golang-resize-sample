package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"

	// "github.com/nfnt/resize"
	"golang.org/x/image/draw"
)

/*
 * 必要に応じてパッケージのコメントアウトを解除
 */

// func Resize() {
// 	img, err := os.Open("./tmp/sample.jpg")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer img.Close()

// 	decordImg, err := jpeg.Decode(img)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	bound := decordImg.Bounds()
// 	fmt.Printf("sample.jpg size \nX: %dpx, Y: %dpx\n", bound.Dx(), bound.Dy())

// 	m := resize.Resize(uint(bound.Dx()/4), uint(bound.Dy()/4), decordImg, resize.Lanczos3)

// 	dst, err := os.Create("./img/resize_sample.jpg")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer dst.Close()

// 	jpeg.Encode(dst, m, nil)
// }

func drawSample() {
	src, err := os.Open("./tmp/sample.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	imgSrc, t, err := image.Decode(src)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("Type of image: ", t)

	rctSrc := imgSrc.Bounds()

	imgDst := image.NewRGBA(image.Rect(0, 0, rctSrc.Dx()/4, rctSrc.Dy()/4)) // 669KB -> 142KB
	draw.CatmullRom.Scale(imgDst, imgDst.Bounds(), imgSrc, rctSrc, draw.Over, nil)

	dst, err := os.Create("./img/resize_sample.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()

	jpeg.Encode(dst, imgDst, &jpeg.Options{Quality: 100})
}

func main() {
	// Resize()
	drawSample()
}
