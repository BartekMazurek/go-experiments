package main

import (
    "fmt"
    "math"
)

type Point struct {
    A float64
    B float64
}

func (p Point) distance(q Point) float64 {
    return math.Hypot(q.A - p.A, q.B - p.B)
}

func main() {

    p := Point{3, 5}
    q := Point{7, 9}

    // method value
    // take the distance method associated with a specific p instance
    // func(q Point) float64
    DistanceFromP := p.distance
    fmt.Println(DistanceFromP(q))

    // method expression
    // take the distance method as a function, but without assigning it to a specific instance
    // func(Point, Point) float64
    Distance := Point.distance
    fmt.Println(Distance(p, q))
}
