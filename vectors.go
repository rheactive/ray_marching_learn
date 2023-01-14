package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type yds_vec struct {
	x float64
	y float64
	z float64
}

func vec_64_new(x float64, y float64, z float64) yds_vec {
	return yds_vec{
		x: x,
		y: y,
		z: z,
	}
}

func vec_64_add(vec1 yds_vec, vec2 yds_vec) yds_vec {
	return yds_vec{
		x: vec1.x + vec2.x,
		y: vec1.y + vec2.y,
		z: vec1.z + vec2.z,
	}
}

func vec_64_sub(vec1 yds_vec, vec2 yds_vec) yds_vec {
	return yds_vec{
		x: vec1.x - vec2.x,
		y: vec1.y - vec2.y,
		z: vec1.z - vec2.z,
	}
}

func vec_64_mul(vec yds_vec, mul float64) yds_vec {
	return yds_vec{
		x: vec.x * mul,
		y: vec.y * mul,
		z: vec.z * mul,
	}
}

func vec_64_norm(vec yds_vec) float64 {
	return math.Sqrt(vec.x*vec.x + vec.y*vec.y + vec.z*vec.z)
}

func vec_64_normalize(vec yds_vec) yds_vec {
	return vec_64_mul(vec, 1.0/vec_64_norm(vec))
}

func vec_64_dot(vec1 yds_vec, vec2 yds_vec) float64 {
	return vec1.x*vec2.x + vec1.y*vec2.y + vec1.z*vec2.z
}

func vec_64_dist(vec1 yds_vec, vec2 yds_vec) float64 {
	return vec_64_norm(vec_64_sub(vec1, vec2))
}

func yds_to_rl(vec yds_vec) rl.Vector3 {
	return rl.NewVector3(float32(vec.x), float32(vec.y), float32(vec.z))
}
