// https://www.youtube.com/watch?v=SqrbIlUwR0U&t=3283s
package main

import (
	"fmt"
	"math"
)

func main() {
	circle := Circle{x: 0, y: 0, radius: 10.3}
	rect := Reactangle{width: 9, height: 6.1}
	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rect Area: %f\n", getArea(rect))
}

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Reactangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Reactangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}