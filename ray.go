package main

import (
	"image/color"
	"math"
)

const maxDepth = 10

type Ray struct {
	Origin, Direction Vec3
}

func (r Ray) PointAtParameter(t float64) Vec3 {
	return r.Origin.Add(r.Direction.Mul(t))
}

func (r Ray) Color(world HitableList, depth int) color.NRGBA64 {
	hit, rec := world.Hit(r, 0.001, math.MaxFloat64)
	if hit {
		if depth < maxDepth {
			if toScatter, scattered, attenuation := rec.material.scatter(r, rec); toScatter {
				return attenuation.MulColor(scattered.Color(world, depth+1))
			}
		} else {
			return Vec3{0, 0, 0}.ToNRGBA64()
		}
	}

	unitDirection := r.Direction.UnitVector()
	t := 0.5 * (unitDirection.y + 1.0)

	return Vec3{1.0, 1.0, 1.0}.Mul(1.0 - t).Add(Vec3{0.5, 0.7, 1.0}.Mul(t)).ToNRGBA64()
}
