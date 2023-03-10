package main

const PI = 3.14159

const SCREEN_WIDTH int32 = 2880
const SCREEN_HEIGHT int32 = 1620
const ASP_RATIO = SCREEN_HEIGHT / SCREEN_WIDTH

const TARGET_FPS = 60

const HALF_WIDTH = SCREEN_WIDTH / 2
const HALF_HEIGHT = SCREEN_HEIGHT / 2

const VIEWPORT_DIST float64 = 1
const FOV = PI / 2.5
const HALF_FOV = FOV / 2

const MAX_DIST float64 = VIEWPORT_DIST * 50
const TOL float64 = 1e-2
const MAX_IT int = 100

const r_min float64 = VIEWPORT_DIST
const r_max float64 = MAX_DIST
const r0 float64 = 20 * TOL






