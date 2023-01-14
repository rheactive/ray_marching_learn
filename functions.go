package main

import (
	//"fmt"
	"math"
	"github.com/gen2brain/raylib-go/raylib"
)

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
				intensity += lights[l].intensity * n_dot_l
			}

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

	if intensity > 1 {
		intensity = 1
	}

	return intensity
}

func light_plane(
	point rl.Vector3,
	ray rl.Vector3, 
	pln plane, 
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
				intensity += lights[l].intensity * n_dot_l
			}

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

	if intensity > 1 {
		intensity = 1
	}

	return intensity
}