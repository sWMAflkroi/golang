package oop

import "fmt"

//Position with x and y
type Position struct {
	x int
	y int
}

//GeoObject with position and color
type GeoObject struct {
	pos   Position
	color int
}

//Circle with radius
type Circle struct {
	GeoObject
	radius int
}

//Rectangle with width and height
type Rectangle struct {
	GeoObject
	width  int
	height int
}

//Triangle with three points
type Triangle struct {
	GeoObject
	po1, po2, po3 Position
}

//Painter Interface
type Painter interface {
	Paint()
}

//Paint function with Circle
func (c Circle) Paint() {
	fmt.Printf("Circle with radius=%v at position=%v and color=%v", c.radius, c.pos, c.color)
}

//Paint function with Rectangle
func (r Rectangle) Paint() {
	fmt.Printf("Rectangle with width=%v and height=%v at position=%v with color=%v", r.width, r.height, r.pos, r.color)
}

//Paint function with Triangle
func (t Triangle) Paint() {
	fmt.Printf("Trinangle with point1=%v, point2=%v, point3=%v at position=%v with color=%v", t.po1, t.po2, t.po3, t.pos, t.color)
}

func main() {
	objects := []Painter{
		Circle{GeoObject{Position{1, 2}, 4}, 20},
		Rectangle{GeoObject{Position{1, 2}, 4}, 10, 20},
		Triangle{GeoObject{Position{1, 2}, 4}, Position{5, 10}, Position{6, 11}, Position{7, 12}},
	}
	for _, v := range objects {
		v.Paint()
	}
}
