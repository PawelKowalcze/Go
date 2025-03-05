package main

import "fmt"

type rectangle struct {
	width  int
	height int
}

func (r rectangle) area() int {
	return r.width * r.height
}

func main() {
	r := rectangle{
		width:  10,
		height: 20,
	}

	fmt.Println(r.area())

}
