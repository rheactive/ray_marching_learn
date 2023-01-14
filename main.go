package main

import (
	//"fmt"
	"math"
	"github.com/gen2brain/raylib-go/raylib"
)

func dist_to_sphere(point rl.Vector3, sph sphere) float32 {
	return rl.Vector3Distance(point, sph.center) - sph.radius
}

func dist_to_plane(point rl.Vector3, pln plane) float32 {
	return rl.Vector3DotProduct(pln.normal, rl.Vector3Subtract(point, pln.point))
	
}

func color_transform (col rl.Vector3) rl.Color {
	return rl.ColorFromNormalized(rl.NewVector4(col.X, col.Y, col.Z, 1))
}

func canvas_to_viewport(u int32, v int32, ratio float32) rl.Vector3 {
	return rl.Vector3Normalize(rl.NewVector3(float32(u) * ratio, float32(v) * ratio, VIEWPORT_DIST))
}

func march_ray(
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
		intensity := light_sphere(point, ray, spheres[id], lights)
		color = rl.Vector3Multiply(spheres[id].color , intensity)
	}

	if obj == "plane" {
		intensity := light_plane(point, ray, planes[id], lights)
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
			color = march_ray(origin, ray, spheres, planes, lights)
			rl.ImageDrawPixel(canvas, 
				HALF_WIDTH + u, 
				HALF_HEIGHT - v, 
				color_transform(color))
		}
	}

	rl.ExportImage(*canvas, "render.png")

}