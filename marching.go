package main

func march_ray_camera(
	origin vec_64,
	ray vec_64,
	spheres []sphere,
	cylinders []cylinder,
	planes []plane,
	lights []light_source,
) vec_64 {
	r := r_min
	point := vec_64_add(origin, vec_64_mul(ray, r))
	color := vec_64_new(1, 1, 1)

	closest_d, id, obj := closest_dist(point, spheres, cylinders, planes)

	it := 0
	for r < r_max && closest_d > TOL && it < MAX_IT {
		it++
		r += closest_d
		point = vec_64_add(origin, vec_64_mul(ray, r))
		closest_d, id, obj = closest_dist(point, spheres, cylinders, planes)
	}

	if obj == "sphere" {
		var intensity float64 = 0
		intensity = light_sphere(point, ray, spheres[id], spheres, cylinders, planes, lights)
		color = vec_64_mul(spheres[id].color, intensity)
	}

	if obj == "plane" {
		var intensity float64 = 0
		intensity = light_plane(point, ray, planes[id], spheres, cylinders, planes, lights)
		color = vec_64_mul(planes[id].color, intensity)
	}

	if obj == "cylinder" {
		var intensity float64 = 0
		intensity = light_cylinder(point, ray, cylinders[id], spheres, cylinders, planes, lights)
		color = vec_64_mul(cylinders[id].color, intensity)
	}

	return color
}

func march_ray_shadow(
	point vec_64,
	spheres []sphere,
	cylinders []cylinder,
	planes []plane,
	light light_source,
	ldir vec_64,
) bool {
	shadow := false

	if light.light != "ambient" {
		r := r0
		rm := r_max

		if light.light == "point" {
			rm = vec_64_dist(light.position, point)
		}

		point2 := vec_64_add(point, vec_64_mul(ldir, r))
		closest_d, _, _ := closest_dist(point2, spheres, cylinders, planes)

		it := 0
		for r < rm && closest_d > TOL && it < MAX_IT {
			it++
			r += closest_d
			point2 = vec_64_add(point, vec_64_mul(ldir, r))
			closest_d, _, _ = closest_dist(point2, spheres, cylinders, planes)
		}

		if closest_d <= TOL {
			shadow = true
		}
	}

	return shadow
}