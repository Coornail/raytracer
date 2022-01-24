package main

type HitRecord struct {
	t        float64
	p        Vec3
	normal   Vec3
	material Material
}

type Hitable interface {
	Hit(r Ray, t_min float64, t_max float64) (bool, HitRecord)
}

type HitableList []Hitable

func (hl HitableList) Hit(r Ray, tMin float64, tMax float64) (bool, HitRecord) {
	hitAnything := false
	closestSoFar := tMax
	var rec HitRecord
	for _, h := range hl {
		if hit, tempRec := h.Hit(r, tMin, closestSoFar); hit {
			hitAnything = true
			closestSoFar = tempRec.t
			rec = tempRec
		}
	}
	return hitAnything, rec
}
