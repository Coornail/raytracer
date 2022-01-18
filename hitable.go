package main

type HitRecord struct {
	t float64
	p Vec3
	normal Vec3
}

type Hitable interface {
	Hit(r Ray, t_min float64, t_max float64) (bool, HitRecord)
}


type HitableList []Hitable 

func (hl HitableList) Hit(r Ray, t_min float64, t_max float64) (bool, HitRecord) {
	hit_anything := false
	closest_so_far := t_max
	var rec HitRecord
	for _, h := range hl {
		if hit, tempRec := h.Hit(r, t_min, closest_so_far); hit {
			hit_anything = true
			closest_so_far = tempRec.t
			rec = tempRec
		}
	}
	return hit_anything, rec
}
