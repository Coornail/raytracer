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

func (v3 Vec3) MulColor(v2 color.Color) color.NRGBA64 {
	r, g, b, _ := v2.RGBA()
	return Vec3{v3.x * (float64(r) / 0xffff), v3.y * (float64(g) / 0xffff), v3.z * (float64(b))/0xffff}.ToNRGBA64()
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

func (v3 Vec3) SquaredLength() float64 {
	return v3.x*v3.x + v3.y*v3.y + v3.z*v3.z
}

func (v3 Vec3) SubScalar(n float64) Vec3 {
	return Vec3{
		v3.x - n,
		v3.y - n,
		v3.z - n,
	}
}

func Reflect(v, n Vec3) Vec3 {
	return v.Sub(n.Mul(2 * dot(v, n)))
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
		}.Sub(Vec3{1, 1, 1})

		if p.SquaredLength() < 1 {
			return p
		}
	}
}
