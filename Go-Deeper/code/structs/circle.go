package main

import (
    "math"
)

type Circle struct {
    radius int
}

func (s *Circle) Area() float64 {
    return math.Pi * float64(s.radius) * float64(s.radius)
}
