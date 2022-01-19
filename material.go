package main

type Material interface {
	scatter(Ray, HitRecord) (bool, Ray, Vec3)
}

type Lambertian struct {
	Albedo Vec3
}

func (l Lambertian) scatter(r Ray, rec HitRecord) (bool, Ray, Vec3) {
	target := rec.p.Add(rec.normal).Add(RandomInUnitSphere())
	return true, Ray{rec.p, target.Sub(rec.p)}, l.Albedo
}

type Metal struct {
	Albedo Vec3
	Fuzz float64
}

func (m Metal) scatter(r Ray, rec HitRecord) (bool, Ray, Vec3) {
	reflected := Reflect(r.Direction.UnitVector(), rec.normal)
	scattered := Ray{rec.p, reflected.Add(RandomInUnitSphere().Mul(m.Fuzz))}
	return dot(scattered.Direction, rec.normal) > 0, scattered, m.Albedo
}
