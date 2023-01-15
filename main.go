package main

import (
	"math"
	"github.com/gen2brain/raylib-go/raylib"
)

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

	cylinders := makeSceneCylinders()

	planes := makeScenePlanes()

	for u := -HALF_WIDTH; u < HALF_WIDTH; u++ {
		for v := -HALF_HEIGHT; v < HALF_HEIGHT; v++ {
			ray = canvas_to_viewport(u, v, ratio)
			color = march_ray_camera(origin, ray, spheres, cylinders, planes, lights)
			rl.ImageDrawPixel(canvas,
				HALF_WIDTH+u,
				HALF_HEIGHT-v,
				color_transform(color))
		}
	}

	rl.ExportImage(*canvas, "render.png")

}
