package main

type Material interface {
	scatter(Ray, HitRecord) (bool, Ray, Vec3)
}

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

func (hl HitableList) Hit(r Ray, t_min float64, t_max float64) (bool, HitRecord) {
	hitAnything := false
	closestSoFar := t_max
	var rec HitRecord
	for _, h := range hl {
		if hit, tempRec := h.Hit(r, t_min, closestSoFar); hit {
			hitAnything = true
			closestSoFar = tempRec.t
			rec = tempRec
		}
	}
	return hitAnything, rec
}
