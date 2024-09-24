package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	circumf() float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s *square) area() float64 {
	return s.length * s.length
}

func (s *square) circumf() float64 {
	return  s.length * 4
}

func (c *circle) area() float64  {
	return math.Pi * c.radius * c.radius
}

func (c *circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

func Shapeinfo(s shape) {
	fmt.Printf("Area of shape: %T is %0.2f \n", s, s.area())
	fmt.Printf("Circumference of shape: %T is %0.2f \n", s, s.circumf())
}

func main() {
	mycircle := circle{
		radius: 10,
	}
	mysquare := square{
		length: 10,
	}

	fmt.Println(mycircle.area())
	fmt.Println(mysquare.area())
	(Shapeinfo(&mycircle))
}

