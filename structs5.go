package main

import (
	"fmt"
	"math"
)

type Vector struct {
	X float64
	Y float64
}

func (v Vector) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vector) Add(other *Vector) Vector {
	return &Vector{X: v.X + other.X, Y: v.Y + other.Y}
}

func main() {
	vec := Vector{3.0, 4.0}
	fmt.Println(vec)
	fmt.Println(vec.Magnitude())
	fmt.Println(vec.Add(&Vector{6.0, 4.0}))
}
