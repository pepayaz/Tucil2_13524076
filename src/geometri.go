package main

import "math"

type Vector3 struct {
	X, Y, Z float64
}

type Segitiga struct {
	V1, V2, V3 int
}

type Model struct {
	Vertices []Vector3
	Faces    []Segitiga
}

type BoundingBox struct {
	Center        Vector3
	HalfDimension float64
}

func hitungRootBox(vertices []Vector3) BoundingBox {
	if len(vertices) == 0 {
		return BoundingBox{}
	}

	min := vertices[0]
	max := vertices[0]

	for _, v := range vertices {
		if v.X < min.X {
			min.X = v.X
		}
		if v.Y < min.Y {
			min.Y = v.Y
		}
		if v.Z < min.Z {
			min.Z = v.Z
		}

		if v.X > max.X {
			max.X = v.X
		}
		if v.Y > max.Y {
			max.Y = v.Y
		}
		if v.Z > max.Z {
			max.Z = v.Z
		}
	}

	center := Vector3{
		X: (min.X + max.X) / 2,
		Y: (min.Y + max.Y) / 2,
		Z: (min.Z + max.Z) / 2,
	}

	diffX := max.X - min.X
	diffY := max.Y - min.Y
	diffZ := max.Z - min.Z

	maxDiff := diffX
	if diffY > maxDiff {
		maxDiff = diffY
	}
	if diffZ > maxDiff {
		maxDiff = diffZ
	}

	padding := maxDiff * 0.01

	return BoundingBox{
		Center:        center,
		HalfDimension: (maxDiff / 2) + padding,
	}
}

func Sub(a, b Vector3) Vector3 {
	return Vector3{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func Cross(a, b Vector3) Vector3 {
	return Vector3{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}

func Dot(a, b Vector3) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func Abs(x float64) float64 {
	return math.Abs(x)
}

func Min3(a, b, c float64) float64 {
	return math.Min(math.Min(a, b), c)
}
func Max3(a, b, c float64) float64 {
	return math.Max(math.Max(a, b), c)
}

func checkIntersection(box BoundingBox, v0, v1, v2 Vector3) bool {
	v0 = Sub(v0, box.Center)
	v1 = Sub(v1, box.Center)
	v2 = Sub(v2, box.Center)

	e0 := Sub(v1, v0)
	e1 := Sub(v2, v1)
	e2 := Sub(v0, v2)

	h := box.HalfDimension

	edges := []Vector3{e0, e1, e2}

	for _, e := range edges {
		if !testSumbu(Vector3{0, -e.Z, e.Y}, v0, v1, v2, h) {
			return false
		}
		if !testSumbu(Vector3{e.Z, 0, -e.X}, v0, v1, v2, h) {
			return false
		}
		if !testSumbu(Vector3{-e.Y, e.X, 0}, v0, v1, v2, h) {
			return false
		}
	}

	if Min3(v0.X, v1.X, v2.X) > h || Max3(v0.X, v1.X, v2.X) < -h {
		return false
	}
	if Min3(v0.Y, v1.Y, v2.Y) > h || Max3(v0.Y, v1.Y, v2.Y) < -h {
		return false
	}
	if Min3(v0.Z, v1.Z, v2.Z) > h || Max3(v0.Z, v1.Z, v2.Z) < -h {
		return false
	}

	normal := Cross(e0, e1)
	return testBidang(normal, v0, h)
}

func testSumbu(axis, v0, v1, v2 Vector3, h float64) bool {
	p0 := Dot(v0, axis)
	p1 := Dot(v1, axis)
	p2 := Dot(v2, axis)

	r := h*Abs(axis.X) + h*Abs(axis.Y) + h*Abs(axis.Z)

	minP := Min3(p0, p1, p2)
	maxP := Max3(p0, p1, p2)

	return !(math.Max(-maxP, minP) > r)
}

func testBidang(normal, v0 Vector3, h float64) bool {
	d := Dot(normal, v0)

	r := h*Abs(normal.X) + h*Abs(normal.Y) + h*Abs(normal.Z)

	return Abs(d) <= r
}
