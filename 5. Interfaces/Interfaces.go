package main

import "fmt"

type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * 3.14 * c.radius
}

// Interface is a collection of method signatures that a Type can implement
func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	fmt.Println("Rectangle area:", r.area())
	fmt.Println("Rectangle perimeter:", r.perimeter())
	fmt.Println("Circle area:", c.area())
	fmt.Println("Circle perimeter:", c.perimeter())
}
