package main

/*Арсений решил выложить на офисном полу кафельную мозаику из геометрических фигур. Для этого нужно сперва рассчитать их площади. Кстати, не создать ли для этого специальный интерфейс?
Создайте интерфейс Shape с методом Area() float64, который будет возвращать площадь фигуры. Создайте структуры Circle с полем radius (float64) и Rectangle с полями width (float64) и height (float64), которые будут реализовывать этот интерфейс.*/
import (
	"math"
)

type Shape interface {
	Area() float64
}

func CalculateArea(s Shape) float64 {
	return s.Area()
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}
