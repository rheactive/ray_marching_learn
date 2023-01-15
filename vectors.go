package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type vec_64 struct {
	x float64
	y float64
	z float64
}

func vec_64_new(x float64, y float64, z float64) vec_64 {
	return vec_64{
		x: x,
		y: y,
		z: z,
	}
}

func vec_64_add(vec1 vec_64, vec2 vec_64) vec_64 {
	return vec_64{
		x: vec1.x + vec2.x,
		y: vec1.y + vec2.y,
		z: vec1.z + vec2.z,
	}
}

func vec_64_sub(vec1 vec_64, vec2 vec_64) vec_64 {
	return vec_64{
		x: vec1.x - vec2.x,
		y: vec1.y - vec2.y,
		z: vec1.z - vec2.z,
	}
}

func vec_64_mul(vec vec_64, mul float64) vec_64 {
	return vec_64{
		x: vec.x * mul,
		y: vec.y * mul,
		z: vec.z * mul,
	}
}

func vec_64_norm(vec vec_64) float64 {
	return math.Sqrt(vec.x*vec.x + vec.y*vec.y + vec.z*vec.z)
}

func vec_64_normalize(vec vec_64) vec_64 {
	return vec_64_mul(vec, 1.0/vec_64_norm(vec))
}

func vec_64_dot(vec1 vec_64, vec2 vec_64) float64 {
	return vec1.x*vec2.x + vec1.y*vec2.y + vec1.z*vec2.z
}

func vec_64_dist(vec1 vec_64, vec2 vec_64) float64 {
	return vec_64_norm(vec_64_sub(vec1, vec2))
}

func yds_to_rl(vec vec_64) rl.Vector3 {
	return rl.NewVector3(float32(vec.x), float32(vec.y), float32(vec.z))
}
