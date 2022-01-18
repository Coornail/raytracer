package main

import (
	"image"
	"image/png"
	"log"
	"os"
)

func main() {
	nx := 2048
	ny := 2048

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
			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)
			r := cam.GetRay(u, v)
			color := r.Color(HitableList(world))
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
