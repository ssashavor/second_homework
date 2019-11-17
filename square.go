package main

import "fmt"

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s *Square) End() (p Point) {
	p.x = s.start.x + int(s.a)
	p.y = s.start.y - int(s.a)
	return p
}

func (s *Square) Perimeter() uint {
	return s.a * 4
}

func (s *Square) Area() uint {
	return s.a * s.a
}

func main() {
	s := Square{Point{1, 1}, 5}
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())
}
