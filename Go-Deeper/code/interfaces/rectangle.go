package main

import ()

type Rectangle struct {
    length int
    width int
}

func (r *Rectangle) Area() float64 {
    return float64(r.length * r.width)
}

func (r *Rectangle) IsSqaure() bool {
    if r.length == r.width {
        return true
    } else {
        return false
    }
}
