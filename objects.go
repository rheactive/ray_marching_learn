package main

import (
	//"fmt"
	//"math"
	"github.com/gen2brain/raylib-go/raylib"
)

type light_source struct {
	light string
	intensity float32
	position rl.Vector3
	direction rl.Vector3
}

func new_light_source(light string, intensity float32) *light_source {
	return &light_source{
		light: light,
		intensity: intensity,
		position: rl.NewVector3(0, 0, 0),
		direction: rl.NewVector3(0, 0, 0),
	}
}

type sphere struct {
	center rl.Vector3
	radius float32
	color rl.Vector3
	specularity int
}

type plane struct {
	normal rl.Vector3
	point rl.Vector3
	color rl.Vector3
	specularity int
}

func newPlane (normal rl.Vector3, point rl.Vector3, color rl.Vector3, specularity int) *plane {
	return &plane {
		normal: rl.Vector3Normalize(normal),
		point: point,
		color: color,
		specularity: specularity,
	}
}

func makeSceneSpheres () []sphere {
	spheres := make([]sphere, 3)

	spheres[0] = sphere{
		center: rl.NewVector3(0, -1, 3.5),
		radius: 1,
		color: rl.NewVector3(0, 0.3, 1),
		specularity: 300,
	}

	spheres[1] = sphere{
		center: rl.NewVector3(2.7, 0, 5.5),
		radius: 1,
		color: rl.NewVector3(1, 0.2, 0),
		specularity: 200,
	}

	spheres[2] = sphere{
		center: rl.NewVector3(-2.7, 0, 5.5),
		radius: 1,
		color: rl.NewVector3(1, 1, 0),
		specularity: 700,
	}

	return spheres
}

func makeScenePlanes () []plane {
	planes := make([]plane, 3)

	planes[0] = *newPlane(
		rl.NewVector3(0, 1, 0), 
		rl.NewVector3(0, -1, 0), 
		rl.NewVector3(0.7, 0.7, 0.8),
		100,
	)

	planes[1] = *newPlane(
		rl.NewVector3(1, 0, -1), 
		rl.NewVector3(-7, 0, 7), 
		rl.NewVector3(0.7, 0.7, 0.8),
		100,
	)

	planes[2] = *newPlane(
		rl.NewVector3(-1, 0, -1), 
		rl.NewVector3(7, 0, 7), 
		rl.NewVector3(0.7, 0.7, 0.8),
		100,
	)

	return planes
}

func makeSceneLights () []light_source {
	lights := make([]light_source, 3)

	lights[0] = *new_light_source("ambient", 0.1)

	lights[1] = *new_light_source("point", 0.5)
	lights[1].position = rl.NewVector3(4, 4, -4)

	lights[2] = *new_light_source("directional", 0.4)
	lights[2].direction = rl.Vector3Normalize(rl.NewVector3(1, 1, 0))

	return lights
}