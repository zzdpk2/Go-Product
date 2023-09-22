package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {

	checkSum := func(t *testing.T, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Test Perimeter function 1", func(t *testing.T) {
		got := Perimeter(10.0, 10.0)
		want := 40.0
		checkSum(t, got, want)
	})

	t.Run("Test Perimeter function 2", func(t *testing.T) {
		rectangle := Rectangle{
			Width:  20.0,
			Height: 22.0,
		}
		got := rectangle.Perimeter()
		want := 84.0
		checkSum(t, got, want)
	})

	t.Run("Test Area 1", func(t *testing.T) {
		got := Perimeter(10.0, 10.0)
		want := 100.0
		checkSum(t, got, want)
	})

	t.Run("Test Area 2", func(t *testing.T) {
		rectangle := Rectangle{
			Width:  12.0,
			Height: 9.0,
		}
		got := AreaByObj(rectangle)
		want := 108.0
		checkSum(t, got, want)
	})
}

func TestArea1(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, &rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, &circle, 314.1592653589793)
	})

}

func TestArea2(t *testing.T) {
	areaTest := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{"Rectangle", &Rectangle{12.0, 6.0}, 72.0},
		{"Circle", &Circle{10.0}, 314.1592653589793},
		{"Triangle", &Triangle{3.0, 4.0, 6.0}, 36.0},
	}

	for _, tt := range areaTest {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("got %.2f want %.2f", got, tt.hasArea)
		}
	}

}
