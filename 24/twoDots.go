package main

import (
	"fmt"
	"math"
)

/*
	TODO:
	Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point{
	return &Point{x, y}
}

func findDistance(a, b *Point) float64 {
	result := math.Sqrt((math.Pow(a.x - b.x, 2) + math.Pow(a.y - b.y, 2)))
	return result
}

func main() {
	a := NewPoint(2, 4)
	b := NewPoint(5, 7)
	fmt.Println(findDistance(a, b))
}