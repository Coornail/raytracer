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
		return Vec3{rec.normal.x + 1, rec.normal.y + 1, rec.normal.z + 1}.Mul(0.5).ToNRGBA64()
	}

	unitDirection := r.Direction.UnitVector()
	t := 0.5 * (unitDirection.y + 1.0)

	return Vec3{1.0, 1.0, 1.0}.Mul(1.0 - t).Add(Vec3{0.5, 0.7, 1.0}.Mul(t)).ToNRGBA64()
}
