package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"
)

func main() {

	a := genMatrix()
	fmt.Println(a)

	width := 100 + 20
	height := 100 + 20

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	cyan := color.RGBA{100, 200, 200, 0xff}

	//Set background
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, color.RGBA{0, 0, 0, 0xff})
		}
	}
	size := 20

	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if a[y][x] == 1 {
				for i := 0; i < size; i++ {
					for j := 0; j < size; j++ {
						img.Set(x*size+j+10, y*size+i+10, cyan)
					}
				}
			}
		}
	}

	// Set color for each pixel.
	/* for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case x < width/2 && y < height/2: // upper left quadrant
				img.Set(x, y, cyan)
			case x >= width/2 && y >= height/2: // lower right quadrant
				img.Set(x, y, color.White)
			default:
				// Use zero value.
			}
		}
	} */

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}

func genMatrix() [5][5]int {
	l := [5][2]int{{}, {}}
	r := [5][2]int{{}, {}}
	m := [5]int{}
	a := [5][5]int{{}, {}}
	rand.Seed(time.Now().UnixNano())
	max := 1

	for y := 0; y < 5; y++ {
		for x := 0; x < 2; x++ {
			temp := rand.Intn(max + 1)
			l[y][x] = temp
			r[y][1-x] = temp
		}
		temp := rand.Intn(max + 1)
		m[y] = temp
	}

	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if x < 2 {
				a[y][x] = l[y][x]
			} else if x == 2 {
				a[y][x] = m[y]
			} else {
				a[y][x] = r[y][x-3]
			}
		}
	}

	return a
}
