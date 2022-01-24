package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"sync"
)

func main() {
	nx := 2048
	ny := 2048
	var antialias uint32 = 16

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [Frame number]\n", os.Args[0])
		os.Exit(1)
	}
	frame, _ := strconv.ParseInt(os.Args[1], 10, 32)

	hitable := []Sphere{
		{
			Center:   Vec3{math.Sin(float64(frame)) / 60, math.Cos(float64(frame))/60 + float64(frame)/60, -1},
			Radius:   0.5,
			Material: Lambertian{Vec3{0.8, 0.3, 0.3}},
		},
		{
			Center:   Vec3{0, -100.5, -1},
			Radius:   100,
			Material: Lambertian{Vec3{0.8, 0.8, 0.0}},
		},
		{
			Center:   Vec3{1, 0, -1},
			Radius:   0.5,
			Material: Metal{Vec3{0.8, 0.6, 0.2}, 0.1},
		},
		{
			Center:   Vec3{-1, 0, -1},
			Radius:   0.5,
			Material: Metal{Vec3{0.8, 0.8, 0.8}, 0.5},
		},
	}

	world := make([]Hitable, len(hitable))
	for i := range hitable {
		world[i] = hitable[i]
	}

	cam := NewCamera()
	img := image.NewNRGBA64(image.Rect(0, 0, nx, ny))
	for j := 0; j < ny; j++ {
		for i := 0; i <= nx; i++ {
			var red, green, blue uint32
			var wg sync.WaitGroup
			wg.Add(int(antialias))
			for s := 0; s < int(antialias); s++ {
				go func(i, j int) {
					u := (float64(i) + rand.Float64()) / float64(nx)
					v := (float64(j) + rand.Float64()) / float64(ny)
					r := cam.GetRay(u, v)
					tmpR, tmpG, tmpB, _ := r.Color(HitableList(world), 0).RGBA()

					red += tmpR
					green += tmpG
					blue += tmpB
					wg.Done()
				}(i, j)
			}
			wg.Wait()

			scaledDownRed := uint16(red / antialias)
			scaledDownGreen := uint16(green / antialias)
			scaledDownBlue := uint16(blue / antialias)
			img.Set(i, j, gammaCorrect(color.NRGBA64{scaledDownRed, scaledDownGreen, scaledDownBlue, 0xffff}))
		}
	}

	os.Mkdir("out", 0777)
	f, err := os.Create(fmt.Sprintf("out/%d.png", frame))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err = png.Encode(f, img); err != nil {
		log.Printf("failed to encode: %v", err)
	}
}
