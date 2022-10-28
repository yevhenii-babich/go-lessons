package main

import "golang.org/x/exp/constraints"

func Min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

func GMin[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func main() {
	println(Min(1.22, 1.23))
	m := GMin[float32](1.22, 1.23)
	println(m)
	m2 := GMin[int](3, 1)
	println(m2)
	// shortened syntax
	println(GMin(1.33, 1.32))
	println(GMin(2, 1))
	println(GMin("2", "1"))
}
