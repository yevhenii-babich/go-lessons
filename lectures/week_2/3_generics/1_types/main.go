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
}
