package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	nx := 2048
	ny := 2048
	var ns uint32 = 4

	hitable := []Sphere{
		{Vec3{0.0, 0.0, -1.0}, 0.5},
		{Vec3{0.0, -100.5, -1.0}, 100},
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
			for s := 0; s < int(ns); s++ {
				u := float64(i) / float64(nx)
				v := float64(j) / float64(ny)
				r := cam.GetRay(u, v)
				tmpR, tmpG, tmpB, _ := r.Color(HitableList(world)).RGBA()

				red += tmpR
				green += tmpG
				blue += tmpB

			}

			scaledDownRed := uint16(red / ns)
			scaledDownGreen := uint16(green /ns)
			scaledDownBlue := uint16(blue / ns)
			img.Set(i, j, gammaCorrect(color.NRGBA64{scaledDownRed, scaledDownGreen, scaledDownBlue, 65535}))
		}
	}

	f, err := os.Create("out.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err = png.Encode(f, img); err != nil {
		log.Printf("failed to encode: %v", err)
	}
}
