package main

import "math"

func HitSphere(center Vec3, radius float64, r Ray) float64 {
	oc := r.Origin.Sub(center)
	a := dot(r.Direction, r.Direction)
	b := 2.0 * dot(oc, r.Direction)
	c := dot(oc, oc) - radius*radius
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return -1.0
	}

	return (-b - math.Sqrt(discriminant)) / (2.0 * a)
}

func dot(v1, v2 Vec3) float64 {
	return v1.x*v2.x + v1.y*v2.y + v1.z*v2.z
}


type Sphere struct {
	Center Vec3
	Radius float64
}

func (s Sphere) Hit(r Ray, t_min float64, t_max float64) (bool, HitRecord) {
	var rec HitRecord
	oc := r.Origin.Sub(s.Center)
	a := dot(r.Direction, r.Direction)
	b := dot(oc, r.Direction)
	c := dot(oc, oc) - (s.Radius*s.Radius)
	discriminant := b*b - a*c
	if discriminant > 0 {
		temp := (-b - math.Sqrt(discriminant)) / a
		if temp < t_max && temp > t_min {
			rec.t = temp
			rec.p = r.PointAtParameter(rec.t)
			rec.normal = rec.p.Sub(s.Center).Div(s.Radius)
			return true, rec
		}

		temp = (-b + math.Sqrt(discriminant)) / a
		if temp < t_max && temp > t_min {
			rec.t = temp
			rec.p = r.PointAtParameter(rec.t)
			rec.normal = rec.p.Sub(s.Center).Div(s.Radius)
			return true, rec
		}
	}

	return false, rec
}
