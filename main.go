package main

import (
	//"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func march_ray_camera(
	origin yds_vec,
	ray yds_vec,
	spheres []sphere,
	planes []plane,
	lights []light_source,
) yds_vec {
	r := r_min
	point := vec_64_add(origin, vec_64_mul(ray, r))
	color := vec_64_new(1, 1, 1)

	closest_d, id, obj := closest_dist(point, spheres, planes)

	it := 0
	for r < r_max && closest_d > TOL && it < MAX_IT {
		it++
		r += closest_d
		point = vec_64_add(origin, vec_64_mul(ray, r))
		closest_d, id, obj = closest_dist(point, spheres, planes)
	}

	if obj == "sphere" {
		var intensity float64 = 0
		intensity = light_sphere(point, ray, spheres[id], spheres, planes, lights)
		color = vec_64_mul(spheres[id].color, intensity)
	}

	if obj == "plane" {
		var intensity float64 = 0
		intensity = light_plane(point, ray, planes[id], spheres, planes, lights)
		color = vec_64_mul(planes[id].color, intensity)
	}

	return color
}

func main() {

	viewport_width := 2 * VIEWPORT_DIST * math.Tan(HALF_FOV)
	ratio := viewport_width / float64(SCREEN_WIDTH)
	origin := vec_64_new(0, 0, 0)
	ray := vec_64_new(0, 0, 0)
	canvas := rl.GenImageColor(int(SCREEN_WIDTH), int(SCREEN_HEIGHT), rl.Black)
	color := vec_64_new(1, 1, 1)

	rl.ExportImage(*canvas, "render_blank.png")

	lights := makeSceneLights()

	spheres := makeSceneSpheres()

	planes := makeScenePlanes()

	for u := -HALF_WIDTH; u < HALF_WIDTH; u++ {
		for v := -HALF_HEIGHT; v < HALF_HEIGHT; v++ {
			ray = canvas_to_viewport(u, v, ratio)
			color = march_ray_camera(origin, ray, spheres, planes, lights)
			rl.ImageDrawPixel(canvas,
				HALF_WIDTH+u,
				HALF_HEIGHT-v,
				color_transform(color))
		}
	}

	rl.ExportImage(*canvas, "render.png")

}
