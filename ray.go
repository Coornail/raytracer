package main

import (
	"image/color"
	"math"
)

type Ray struct {
	Origin, Direction Vec3
}

func (r Ray) PointAtParameter(t float64) Vec3 {
	return r.Origin.Add(r.Direction.Mul(t))
}

func (r Ray) Color(world HitableList) color.NRGBA64 {
	hit, rec := world.Hit(r, 0.001, math.MaxFloat64)
	if hit {
		target := rec.p.Add(rec.normal).Add(RandomInUnitSphere())
		c := Ray{rec.p, target.Sub(rec.p)}.Color(world)
		r, g, b, _ := c.RGBA()
		return color.NRGBA64{uint16(r/2), uint16(g/2), uint16(b/2), 0xffff}
	}

	unitDirection := r.Direction.UnitVector()
	t := 0.5 * (unitDirection.y + 1.0)

	return Vec3{1.0, 1.0, 1.0}.Mul(1.0 - t).Add(Vec3{0.5, 0.7, 1.0}.Mul(t)).ToNRGBA64()
}
