package main

import (
	"image/color"
	"math"
	"math/rand"
)

type Vec3 struct {
	x, y, z float64
}

func (v3 Vec3) Add(v2 Vec3) Vec3 {
	return Vec3{v3.x + v2.x, v3.y + v2.y, v3.z + v2.z}
}

func (v3 Vec3) Sub(v2 Vec3) Vec3 {
	return Vec3{v3.x - v2.x, v3.y - v2.y, v3.z - v2.z}
}

func (v3 Vec3) Mul(f float64) Vec3 {
	return Vec3{v3.x * f, v3.y * f, v3.z * f}
}

func (v3 Vec3) Div(f float64) Vec3 {
	return Vec3{v3.x / f, v3.y / f, v3.z / f}
}

func (v3 Vec3) UnitVector() Vec3 {
	return v3.Mul(1 / v3.Length())
}

func (v3 Vec3) Length() float64 {
	return math.Sqrt(v3.x*v3.x + v3.y*v3.y + v3.z*v3.z)
}

func (v3 Vec3) ToNRGBA64() color.NRGBA64 {
	return color.NRGBA64{uint16(v3.x * 65535), uint16(v3.y * 65535), uint16(v3.z * 65535), 65535}
}

func gammaCorrect(c color.NRGBA64) color.NRGBA64 {
	return color.NRGBA64{
		uint16(math.Pow((float64(c.R))/65535, 2.2) * 65535),
		uint16(math.Pow((float64(c.G))/65535, 2.2) * 65535),
		uint16(math.Pow((float64(c.B))/65535, 2.2) * 65535),
		c.A,
	}
}

func RandomInUnitSphere() Vec3 {
	for {
		p := Vec3{
			2*rand.Float64() - 1,
			2*rand.Float64() - 1,
			2*rand.Float64() - 1,
		}

		if p.Length() < 1 {
			return p
		}
	}
}
