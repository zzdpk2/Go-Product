package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Shape interface {
	Perimeter() float64
	Area() float64
}

func Perimeter(width float64, height float64) float64 {
	return 2 * (width + height)
}

func Area(width float64, height float64) float64 {
	return width * height
}

func AreaByObj(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

func (rectangle *Rectangle) Perimeter() float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (rectangle *Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

type Circle struct {
	radius float64
}

func (circle *Circle) Perimeter() float64 {
	return 2 * math.Pi * circle.radius
}

func (circle *Circle) Area() float64 {
	return math.Pi * circle.radius * circle.radius
}

type Triangle struct {
	Side1 float64
	Side2 float64
	Side3 float64
}

func (t *Triangle) Area() float64 {
	s := t.Perimeter() * 0.5
	return math.Sqrt(s * (s - t.Side1) * (s - t.Side2) * (s - t.Side3))
}

func (t *Triangle) Perimeter() float64 {
	return t.Side1 + t.Side2 + t.Side3
}
