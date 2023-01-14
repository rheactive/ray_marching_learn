package main

import (
	//"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func dist_to_sphere(point yds_vec, sph sphere) float64 {
	dist := vec_64_dist(point, sph.center) - sph.radius
	return dist
}

func dist_to_plane(point yds_vec, pln plane) float64 {
	dist := vec_64_dot(pln.normal, vec_64_sub(point, pln.point))
	return dist
}

func color_transform(col yds_vec) rl.Color {
	return rl.ColorFromNormalized(rl.NewVector4(float32(col.x), float32(col.y), float32(col.z), 1))
}

func canvas_to_viewport(u int32, v int32, ratio float64) yds_vec {
	return vec_64_normalize(vec_64_new(float64(u)*ratio, float64(v)*ratio, VIEWPORT_DIST))
}

func closest_dist(point yds_vec, spheres []sphere, planes []plane) (float64, int, string) {
	sphere_num := len(spheres)
	plane_num := len(planes)
	closest_dist := MAX_DIST
	obj := "space"
	var id int
	for s := 0; s < sphere_num; s++ {
		d := dist_to_sphere(point, spheres[s])
		if d < closest_dist {
			closest_dist = d
			id = s
			obj = "sphere"
		}
	}
	for p := 0; p < plane_num; p++ {
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
	point yds_vec,
	ray yds_vec,
	sph sphere,
	spheres []sphere,
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
				if !march_ray_shadow(point, spheres, planes, lights[l], ldir) {
					intensity += lights[l].intensity * n_dot_l
					ref := vec_64_sub(
						vec_64_mul(normal, 2*vec_64_dot(normal, ldir)), ldir)
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
	point yds_vec,
	ray yds_vec,
	pln plane,
	spheres []sphere,
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
				if !march_ray_shadow(point, spheres, planes, lights[l], ldir) {
					intensity += lights[l].intensity * n_dot_l
					ref := vec_64_sub(
						vec_64_mul(normal, 2*vec_64_dot(normal, ldir)), ldir)
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

func march_ray_shadow(
	point yds_vec,
	spheres []sphere,
	planes []plane,
	light light_source,
	ldir yds_vec,
) bool {
	shadow := false

	if light.light != "ambient" {
		r := r0
		rm := r_max

		if light.light == "point" {
			rm = vec_64_dist(light.position, point)
		}

		point2 := vec_64_add(point, vec_64_mul(ldir, r))
		closest_d, _, _ := closest_dist(point2, spheres, planes)

		it := 0
		for r < rm && closest_d > TOL && it < MAX_IT {
			it++
			r += closest_d
			point2 = vec_64_add(point, vec_64_mul(ldir, r))
			closest_d, _, _ = closest_dist(point2, spheres, planes)
		}

		if closest_d <= TOL {
			shadow = true
		}
	}

	return shadow
}
