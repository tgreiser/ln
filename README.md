# 3D Line Art Renderer

`ln` is a vector-based 3D rendering engine written in Go. It is used to produce
2-dimensional vector graphics depicting 3-dimensional scenes.

## Features

- Primitives
	- Sphere
	- Cube
	- Triangle
	- Cylinder
- 3D Functions
- Triangle Meshes
	- OBJ & STL
- Vector-based "Texturing"
- CSG (Constructive Solid Geometry) Operations
	- Intersection
	- Difference
	- Union
- Output to PNG or SVG

## How it Works

To understand how `ln` works, it's useful to start with the `Shape` interface:

```go
type Shape interface {
	Paths() Paths
	Intersect(Ray) Hit
	Contains(Vector, float64) bool
	BoundingBox() Box
	Compile()
}
```

Each shape must provide some `Paths` which are 3D polylines on the surface
of the solid. Ultimately anything drawn in the final image is based on these
paths. These paths can be anything. For a sphere they could be lat/lng grid
lines, a triangulated-looking surface, dots on the surface, etc. This is what
we call vector-based texturing.

Each shape must also provide an `Intersect` method that lets the engine test
for ray-solid intersection. This is how the engine knows what is visible to the
camera and what is hidden.

All of the `Paths` are chopped up to some granularity and each point is tested
by shooting a ray toward the camera. If there is no intersection, that point is
visible. If there is an intersection, it is hidden and will not be rendered.

The visible points are then transformed into 2D space using transformation
matrices. The result can then be rendered as PNG or SVG.

The `Contains` method is only needed for CSG (Constructive Solid Geometry)
operations.

## Hello World: A Single Cube

### The Code

```go
package main

import "github.com/fogleman/ln/ln"

func main() {
	// create a scene and add a single cube
	scene := ln.Scene{}
	scene.Add(ln.NewCube(ln.Vector{-1, -1, -1}, ln.Vector{1, 1, 1}))

	// define camera parameters
	eye := ln.Vector{4, 3, 2}    // camera position
	center := ln.Vector{0, 0, 0} // camera looks at
	up := ln.Vector{0, 0, 1}     // up direction

	// define rendering parameters
	width := 1024.0  // rendered width
	height := 1024.0 // rendered height
	fovy := 50.0     // vertical field of view, degrees
	znear := 0.1     // near z plane
	zfar := 10.0     // far z plane
	step := 0.01     // how finely to chop the paths for visibility testing

	// compute 2D paths that depict the 3D scene
	paths := scene.Render(eye, center, up, width, height, fovy, znear, zfar, step)

	// render the paths in an image
	paths.WriteToPNG("out.png", width, height)

	// save the paths as an svg
	paths.WriteToSVG("out.svg", width, height)
}
```

### The Output

![Cube](http://i.imgur.com/d2dGrOJ.png)