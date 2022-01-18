package main

type Camera struct {
	Origin, LowerLeftCorner, Horizontal, Vertical Vec3
}

func (c Camera) GetRay(u, v float64) Ray {
	return Ray{
		Origin: c.Origin,
		Direction: c.LowerLeftCorner.Add(c.Horizontal.Mul(u)).Add(c.Vertical.Mul(v)),
	}
}

func NewCamera() *Camera {
	return &Camera{
		LowerLeftCorner: Vec3{-2.0, -1.0, -1.0},
		Horizontal:     Vec3{4.0, 0.0, 0.0},
		Vertical:      Vec3{0.0, 4.0, 0.0},
		Origin: 	Vec3{0.0, 0.0, 0.0},
	}
}
