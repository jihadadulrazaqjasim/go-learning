package main

import "fmt"

type shape interface {
	area() int
}

type square struct {
	sideLength int
}

type rectangular struct {
	length int
	width  int
}

func (r rectangular) area() int {
	return r.length * r.width
}

func (s square) area() int {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Printf("Type: %T, Size of the shape is: %d\n", s, s.area())
}
