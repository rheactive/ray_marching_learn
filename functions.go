package main

import (
	//"fmt"
	"math"
	"github.com/gen2brain/raylib-go/raylib"
)

func dist_to_sphere(point rl.Vector3, sph sphere) float32 {
	dist := rl.Vector3Distance(point, sph.center) - sph.radius
	return dist
}

func dist_to_plane(point rl.Vector3, pln plane) float32 {
	dist := rl.Vector3DotProduct(pln.normal, rl.Vector3Subtract(point, pln.point))
	return dist
}

func color_transform (col rl.Vector3) rl.Color {
	return rl.ColorFromNormalized(rl.NewVector4(col.X, col.Y, col.Z, 1))
}

func canvas_to_viewport(u int32, v int32, ratio float32) rl.Vector3 {
	return rl.Vector3Normalize(rl.NewVector3(float32(u) * ratio, float32(v) * ratio, VIEWPORT_DIST))
}

func closest_dist(point rl.Vector3, spheres []sphere, planes []plane) (float32, int, string) {
	sphere_num := len(spheres)
	plane_num := len(planes)
	closest_dist := MAX_DIST;
	obj := "space"
	var id int;
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
	point rl.Vector3,
	ray rl.Vector3,
	sph sphere,
	spheres []sphere,
	planes []plane,
	lights []light_source) float32 {
	var intensity float32 = 0

	normal := rl.Vector3Normalize(rl.Vector3Subtract(point, sph.center))
	ldir := rl.NewVector3(0, 0, 0)
	var n_dot_l float32 = 1

	for l := 0; l < len(lights); l++ {
		if lights[l].light == "ambient" {
			intensity += lights[l].intensity
		} else {
			if lights[l].light == "point" {
				ldir = rl.Vector3Normalize(rl.Vector3Subtract(lights[l].position, point)) 
			} 
			if lights[l].light == "directional" {
				ldir = lights[l].direction
			}
			n_dot_l = rl.Vector3DotProduct(normal, ldir)
			if n_dot_l > 0 {
				if !march_ray_shadow(point, spheres, planes, lights[l], ldir) {
					intensity += lights[l].intensity * n_dot_l
					ref := rl.Vector3Subtract(
						rl.Vector3Multiply(normal, 2 * rl.Vector3DotProduct(normal, ldir)), ldir)
					ref_dot_ray := rl.Vector3DotProduct(ref, rl.Vector3Negate(ray))
					
					if ref_dot_ray > 0 {
						intensity += lights[l].intensity * float32(
							math.Pow(float64(ref_dot_ray), 
							float64(sph.specularity)))
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
	point rl.Vector3,
	ray rl.Vector3, 
	pln plane, 
	spheres []sphere,
	planes []plane,
	lights []light_source) float32 {
	var intensity float32 = 0

	normal := pln.normal
	ldir := normal
	var n_dot_l float32 = 1

	for l := 0; l < len(lights); l++ {
		if lights[l].light == "ambient" {
			intensity += lights[l].intensity
		} else {
			if lights[l].light == "point" {
				ldir = rl.Vector3Normalize(rl.Vector3Subtract(lights[l].position, point)) 
			} 
			if lights[l].light == "directional" {
				ldir = lights[l].direction
			}
			n_dot_l = rl.Vector3DotProduct(normal, ldir)
			if n_dot_l > 0 {
				if !march_ray_shadow(point, spheres, planes, lights[l], ldir) {
					intensity += lights[l].intensity * n_dot_l
					ref := rl.Vector3Subtract(
						rl.Vector3Multiply(normal, 2 * rl.Vector3DotProduct(normal, ldir)), ldir)
					ref_dot_ray := rl.Vector3DotProduct(ref, rl.Vector3Negate(ray))
					
					if ref_dot_ray > 0 {
						intensity += lights[l].intensity * float32(
							math.Pow(float64(ref_dot_ray), 
							float64(pln.specularity)))
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
	point rl.Vector3, 
	spheres []sphere,
	planes []plane,
	light light_source,
	ldir rl.Vector3, 
	) bool {
	shadow := false

	if light.light != "ambient" {
		r := 20 * TOL
		rm := r_max

		if light.light == "point" {
			rm = rl.Vector3Distance(light.position, point)
		}

			point2 := rl.Vector3Add(point, rl.Vector3Multiply(ldir, r))
			closest_d, _, obj := closest_dist(point2, spheres, planes)

		it := 0
		for r < rm && closest_d > TOL && it < MAX_IT {
			it ++
			if closest_d < 0 {
				shadow = true
				break
			} else {
				r += closest_d
				point2 = rl.Vector3Add(point2, rl.Vector3Multiply(ldir, r))
				closest_d, _, obj = closest_dist(point2, spheres, planes)
			}
		}

		if obj == "sphere" {
			shadow = true
		}
	}
	
	return shadow
}
