package main

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

type sphere struct {
	center      vec_64
	radius      float64
	color       vec_64
	specularity int
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

func makeSceneSpheres() []sphere {
	spheres := make([]sphere, 3)

	spheres[0] = sphere{
		center:      vec_64_new(0, -1, 3.5),
		radius:      1,
		color:       vec_64_new(0.5, 0.3, 1),
		specularity: 300,
	}

	spheres[1] = sphere{
		center:      vec_64_new(2.7, 0, 5.5),
		radius:      1,
		color:       vec_64_new(1, 0.2, 0),
		specularity: 200,
	}

	spheres[2] = sphere{
		center:      vec_64_new(-2, 0, 6),
		radius:      1.3,
		color:       vec_64_new(1, 1, 0),
		specularity: 700,
	}

	return spheres
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

func makeSceneLights() []light_source {
	lights := make([]light_source, 3)

	lights[0] = *new_light_source("ambient", 0.2)

	lights[1] = *new_light_source("directional", 0.4)
	lights[1].direction = vec_64_normalize(vec_64_new(1, 0.5, -1))

	lights[2] = *new_light_source("point", 0.4)
	lights[2].position = vec_64_new(4, 3, 2)

	return lights
}
