package main

import (
	"math"
	"github.com/gen2brain/raylib-go/raylib"
)

func color_transform(col vec_64) rl.Color {
	return rl.ColorFromNormalized(rl.NewVector4(float32(col.x), float32(col.y), float32(col.z), 1))
}

func canvas_to_viewport(u int32, v int32, ratio float64) vec_64 {
	return vec_64_normalize(vec_64_new(float64(u)*ratio, float64(v)*ratio, VIEWPORT_DIST))
}

func closest_dist(point vec_64, spheres []sphere, cylinders []cylinder, planes []plane) (float64, int, string) {
	closest_dist := MAX_DIST
	obj := "space"
	var id int
	for s := 0; s < len(spheres); s++ {
		d := dist_to_sphere(point, spheres[s])
		if d < closest_dist {
			closest_dist = d
			id = s
			obj = "sphere"
		}
	}
	for c := 0; c < len(cylinders); c++ {
		d := dist_to_cylinder(point, cylinders[c])
		if d < closest_dist {
			closest_dist = d
			id = c
			obj = "cylinder"
		}
	}
	for p := 0; p < len(planes); p++ {
		d := dist_to_plane(point, planes[p])
		if d < closest_dist {
			closest_dist = d
			id = p
			obj = "plane"
		}
	}
	return closest_dist, id, obj
}

func light_sphere(
	point vec_64,
	ray vec_64,
	sph sphere,
	spheres []sphere,
	cylinders []cylinder,
	planes []plane,
	lights []light_source) float64 {
	var intensity float64 = 0

	normal := vec_64_normalize(vec_64_sub(point, sph.center))
	ldir := vec_64_new(0, 0, 0)
	var n_dot_l float64 = 1

	for l := 0; l < len(lights); l++ {
		if lights[l].light == "ambient" {
			intensity += lights[l].intensity
		} else {
			if lights[l].light == "point" {
				ldir = vec_64_normalize(vec_64_sub(lights[l].position, point))
			}
			if lights[l].light == "directional" {
				ldir = lights[l].direction
			}
			n_dot_l = vec_64_dot(normal, ldir)
			if n_dot_l > 0 {
				if !march_ray_shadow(point, spheres, cylinders, planes, lights[l], ldir) {
					intensity += lights[l].intensity * n_dot_l
					ref := vec_64_ref(normal, ldir)
					ref_dot_ray := vec_64_dot(ref, vec_64_mul(ray, -1))

					if ref_dot_ray > 0 {
						intensity += lights[l].intensity *
							math.Pow(ref_dot_ray,
								float64(sph.specularity))
					}
				}

			}

		}

	}

	if intensity > 1 {
		intensity = 1
	}

	return intensity
}

func light_plane(
	point vec_64,
	ray vec_64,
	pln plane,
	spheres []sphere,
	cylinders []cylinder,
	planes []plane,
	lights []light_source) float64 {
	var intensity float64 = 0

	normal := pln.normal
	ldir := normal
	var n_dot_l float64 = 1

	for l := 0; l < len(lights); l++ {
		if lights[l].light == "ambient" {
			intensity += lights[l].intensity
		} else {
			if lights[l].light == "point" {
				ldir = vec_64_normalize(vec_64_sub(lights[l].position, point))
			}
			if lights[l].light == "directional" {
				ldir = lights[l].direction
			}
			n_dot_l = vec_64_dot(normal, ldir)
			if n_dot_l > 0 {
				if !march_ray_shadow(point, spheres, cylinders, planes, lights[l], ldir) {
					intensity += lights[l].intensity * n_dot_l
					ref := vec_64_ref(normal, ldir)
					ref_dot_ray := vec_64_dot(ref, vec_64_mul(ray, -1))

					if ref_dot_ray > 0 {
						intensity += lights[l].intensity *
							math.Pow(ref_dot_ray,
								float64(pln.specularity))
					}
				}

			}

		}

	}

	if intensity > 1 {
		intensity = 1
	}

	return intensity
}

func light_cylinder(
	point vec_64,
	ray vec_64,
	cyl cylinder,
	spheres []sphere,
	cylinders []cylinder,
	planes []plane,
	lights []light_source) float64 {
	var intensity float64 = 0

	normal := cylinder_normal(point, cyl)
	ldir := vec_64_new(0, 0, 0)
	var n_dot_l float64 = 1

	for l := 0; l < len(lights); l++ {
		if lights[l].light == "ambient" {
			intensity += lights[l].intensity
		} else {
			if lights[l].light == "point" {
				ldir = vec_64_normalize(vec_64_sub(lights[l].position, point))
			}
			if lights[l].light == "directional" {
				ldir = lights[l].direction
			}
			n_dot_l = vec_64_dot(normal, ldir)
			if n_dot_l > 0 {
				if !march_ray_shadow(point, spheres, cylinders, planes, lights[l], ldir) {
					intensity += lights[l].intensity * n_dot_l
					ref := vec_64_ref(normal, ldir)
					ref_dot_ray := vec_64_dot(ref, vec_64_mul(ray, -1))

					if ref_dot_ray > 0 {
						intensity += lights[l].intensity *
							math.Pow(ref_dot_ray,
								float64(cyl.specularity))
					}
				}

			}

		}

	}

	if intensity > 1 {
		intensity = 1
	}

	return intensity
}