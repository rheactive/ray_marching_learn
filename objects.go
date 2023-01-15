package main

import "math"

type light_source struct {
	light     string
	intensity float64
	position  vec_64
	direction vec_64
}

func new_light_source(light string, intensity float64) *light_source {
	return &light_source{
		light:     light,
		intensity: intensity,
		position:  vec_64_new(0, 0, 0),
		direction: vec_64_new(0, 0, 0),
	}
}

func makeSceneLights() []light_source {
	lights := make([]light_source, 3)

	lights[0] = *new_light_source("ambient", 0.2)

	lights[1] = *new_light_source("directional", 0.4)
	lights[1].direction = vec_64_normalize(vec_64_new(1, 0.5, -1))

	lights[2] = *new_light_source("point", 0.4)
	lights[2].position = vec_64_new(4, 3, 2)

	return lights
}


type sphere struct {
	center      vec_64
	radius      float64
	color       vec_64
	specularity int
}

func makeSceneSpheres() []sphere {
	spheres := make([]sphere,1)

	spheres[0] = sphere{
		center:      vec_64_new(1.5, 0, 4),
		radius:      1,
		color:       vec_64_new(1, 1, 0),
		specularity: 700,
	}

	return spheres
}

type plane struct {
	normal      vec_64
	point       vec_64
	color       vec_64
	specularity int
}

func newPlane(normal vec_64, point vec_64, color vec_64, specularity int) *plane {
	return &plane{
		normal:      vec_64_normalize(normal),
		point:       point,
		color:       color,
		specularity: specularity,
	}
}

func makeScenePlanes() []plane {
	planes := make([]plane, 3)

	planes[0] = *newPlane(
		vec_64_new(0, 1, 0),
		vec_64_new(0, -1, 0),
		vec_64_new(0.7, 0.7, 0.8),
		100,
	)

	planes[1] = *newPlane(
		vec_64_new(1, 0, -1),
		vec_64_new(-7, 0, 7),
		vec_64_new(0.7, 0.7, 0.8),
		100,
	)

	planes[2] = *newPlane(
		vec_64_new(-1, 0, -1),
		vec_64_new(7, 0, 7),
		vec_64_new(0.7, 0.7, 0.8),
		100,
	)

	return planes
}

type cylinder struct {
	point1      vec_64
	point2      vec_64
	height		float64
	radius      float64
	sphere1		sphere
	sphere2		sphere
	color       vec_64
	specularity int
}

func newCylinder(point1 vec_64, point2 vec_64, radius float64, color vec_64, specularity int) *cylinder {
	return &cylinder{
		point1:		point1,
		point2:		point2,
		height:		vec_64_dist(point2, point1),
		radius:		radius,
		sphere1:	sphere{
			center:      point1,
			radius:      radius,
			color:       color,
			specularity: specularity,
		},
		sphere2:	sphere{
			center:      point2,
			radius:      radius,
			color:       color,
			specularity: specularity,
		},
		color:		color,
		specularity:	specularity,
	}
}

func makeSceneCylinders() []cylinder {
	cylinders := make([]cylinder, 1)

	cylinders[0] = *newCylinder(
		vec_64_new(0, -0.25, 5),
		vec_64_new(-1.8, -0.25, 4),
		0.75,
		vec_64_new(0.5, 0.3, 1),
		1000,
	)

	return cylinders
}

func dist_to_sphere(point vec_64, sph sphere) float64 {
	dist := vec_64_dist(point, sph.center) - sph.radius
	return dist
}

func dist_to_plane(point vec_64, pln plane) float64 {
	dist := vec_64_dot(pln.normal, vec_64_sub(point, pln.point))
	return dist
}

func cylinder_axis_point(point vec_64, cyl cylinder) vec_64 {
	vec1 := vec_64_mul(vec_64_sub(cyl.point2, cyl.point1), 1.0 / cyl.height)
	vec2 := vec_64_sub(point, cyl.point1)
	dot := vec_64_dot(vec1, vec2)
	vec3 := vec_64_add(cyl.point1, vec_64_mul(vec1, dot))
	return vec3
}

func dist_to_cylinder_1(point vec_64, cyl cylinder) float64 {
	vec := cylinder_axis_point(point, cyl)
	dist := vec_64_dist(point, vec)- cyl.radius
	return dist
}

func dist_to_cylinder(point vec_64, cyl cylinder) float64 {
	vec := cylinder_axis_point(point, cyl)
	dist := vec_64_dist(point, vec)- cyl.radius
	if vec_64_dist(vec, cyl.point1) > cyl.height || vec_64_dist(vec, cyl.point2) > cyl.height {
		d1 := dist_to_sphere(point, cyl.sphere1)
		d2 := dist_to_sphere(point, cyl.sphere2)
		dist = math.Min(d1, d2)
	}
	return dist
}

func cylinder_normal(point vec_64, cyl cylinder) vec_64 {
	vec := cylinder_axis_point(point, cyl)
	normal := vec_64_normalize(vec_64_sub(point,vec))
	if vec_64_dist(vec, cyl.point1) > cyl.height || vec_64_dist(vec, cyl.point2) > cyl.height {
		d1 := vec_64_dist(point,cyl.point1)
		d2 := vec_64_dist(point,cyl.point2)
		if d1 < d2 {
			normal = vec_64_normalize(vec_64_sub(point,cyl.point1))
		} else {
			normal = vec_64_normalize(vec_64_sub(point,cyl.point2))
		}
	}
	return normal
}