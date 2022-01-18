package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
)

func main() {
	nx := 2048
	ny := 2048

	var lowerLeftCorner Vec3 = Vec3{-2.0, -1.0, -1.0}
	var vertical Vec3 = Vec3{4.0, 0.0, 0.0}
	var horizontal Vec3 = Vec3{0.0, 4.0, 0.0}
	var origin Vec3 = Vec3{0.0, 0.0, 0.0}

	hitable := []Sphere{
		{Vec3{0.0, 0.0, -1.0}, 0.5},
		{Vec3{0.0, -100.5, -1.0}, 100},
	}

	world := make([]Hitable, len(hitable))
	for i := range hitable {
		world[i] = hitable[i]
	}

	img := image.NewNRGBA64(image.Rect(0, 0, nx, ny))
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i <= nx; i++ {
			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)
			r := Ray{origin, lowerLeftCorner.Add(horizontal.Mul(u)).Add(vertical.Mul(v))}
			color := r.Color(HitableList(world))
			fmt.Printf("%d %d %d\n", color.R, color.G, color.B)
			img.Set(i, j, color)
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
