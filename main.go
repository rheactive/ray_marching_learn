package main

import (
	//"fmt"
	"math"
	"github.com/gen2brain/raylib-go/raylib"
)

func march_ray_camera(
	origin rl.Vector3, 
	ray rl.Vector3, 
	spheres []sphere,
	planes []plane, 
	lights []light_source,
	) rl.Vector3 {
	r := r_min
	point := rl.Vector3Add(origin, rl.Vector3Multiply(ray, r))
	color := rl.NewVector3(1, 1, 1)
	
	closest_d, id, obj := closest_dist(point, spheres, planes)

	it := 0
	for r < r_max && closest_d > TOL && it < MAX_IT {
		it ++
		r += closest_d
		point = rl.Vector3Add(origin, rl.Vector3Multiply(ray, r))
		closest_d, id, obj = closest_dist(point, spheres, planes)
	}

	if obj == "sphere" {
		var intensity float32 = 0 
		intensity = light_sphere(point, ray, spheres[id], spheres, planes, lights)
		color = rl.Vector3Multiply(spheres[id].color, intensity)
	}

	if obj == "plane" {
		var intensity float32 = 0 
		intensity = light_plane(point, ray, planes[id], spheres, planes, lights)
		color = rl.Vector3Multiply(planes[id].color , intensity)
	}
		

	return color
}

func main() {

	viewport_width := 2 * VIEWPORT_DIST * float32(math.Tan(HALF_FOV))
	ratio := viewport_width / float32(SCREEN_WIDTH)
	origin := rl.NewVector3(0, 0, 0)
	ray := rl.NewVector3(0, 0, 0)
	canvas := rl.GenImageColor(int(SCREEN_WIDTH), int(SCREEN_HEIGHT), rl.Black)
	color := rl.NewVector3(1, 1, 1)

	rl.ExportImage(*canvas, "render_blank.png")

	lights := makeSceneLights ()

	spheres := makeSceneSpheres()

	planes := makeScenePlanes()

	for u := -HALF_WIDTH; u < HALF_WIDTH; u++ {
		for v := -HALF_HEIGHT; v < HALF_HEIGHT; v++ {
			ray = canvas_to_viewport(u, v, ratio)
			color = march_ray_camera(origin, ray, spheres, planes, lights)
			rl.ImageDrawPixel(canvas, 
				HALF_WIDTH + u, 
				HALF_HEIGHT - v, 
				color_transform(color))
		}
	}

	rl.ExportImage(*canvas, "render.png")

}