package main

import (
	"errors"
	"fmt"
)

const pi = 3.1415

func main() {
	printCircleArea(5)
	printRectangle(10, 5)
	printTriangle(10, 5, 15)
}

//Круг
func printCircleArea(radius int) {
	circleArea, err := calculateCircleArea(radius)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Радиус круга: %d см.\n", radius)
	fmt.Println("Формула для рассчета площади круга: A= Пr2\n")
	fmt.Printf("Площадь круга: %F см.кв.\n\n", circleArea)
}
func calculateCircleArea(radius int) (float32, error) {
	if radius <= 0 {
		return float32(0), errors.New("Радиус круга не может быть отрицательным!")
	}
	return float32(radius) * float32(radius) * pi, nil
}

//Прямоугольник(Площадь)
func printRectangle(long, breadth int) {
	RectanglePerimetr, err := CalculateRectanglePerimetr(long, breadth)
	RectangleArea, err := calculateRectangleArea(long, breadth)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Длина прямоугольника: %d см.\n", long)
	fmt.Printf("Ширина прямоугольника:%d см.\n", breadth)
	fmt.Println("Формула подсчета площади прямоугольника: S=a*b\n")
	fmt.Printf("Площадь прямоугольника: %d см. кв.\n", RectangleArea)
	fmt.Printf("Периметр прямоугольника: %d см.\n\n", RectanglePerimetr)

}
func calculateRectangleArea(long, breadth int) (int64, error) {
	RectangleArea := long * breadth
	if RectangleArea <= 0 {
		return int64(0), errors.New("Площaдь не может быть отрицательная! Проверьте правильность введенных данных...")
	}
	return int64(long) * int64(breadth), nil
}
func CalculateRectanglePerimetr(long, breadth int) (int64, error) {
	RectanglePerimetr := (long + breadth) * 2
	if RectanglePerimetr <= 0 {
		return int64(0), errors.New("Периметр не может быть отрицательным, проверьте правильность введенных данных....")
	}
	return (int64(long) * int64(breadth)) * 2, nil
}

//Треугольник (Площадь + Периметр)
func printTriangle(a, b, c int) {
	Perimetr, err := CalculatePerimetr(a, b, c)
	TriangleArea, err := calculateTriangleArea(a, b, c)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Первая сторона:%d см. \n", a)
	fmt.Printf("Вторая сторона:%d см.\n", b)
	fmt.Printf("Третья сторона:%d см.\n", c)
	fmt.Println("Формула подсчета площади треугольника по 3 сторонам:S = a*b*c\n")
	fmt.Printf("Площадь треугольника:%d см.кв.\n", TriangleArea)
	fmt.Printf("Периметр треугольника:%d см.\n", Perimetr)
}
func calculateTriangleArea(a, b, c int) (int64, error) {
	TriangleArea := a * b * c
	if TriangleArea <= 0 {
		return int64(0), errors.New("Площадь быть отрицательной не может! Провверьте правильность введенных данных...")
	}
	return int64(a) * int64(b) * int64(c), nil
}
func CalculatePerimetr(a, b, c int) (int64, error) {
	Perimetr := a + b + c
	if Perimetr <= 0 {
		return int64(0), errors.New("Проверьте правильность набора данных....")
	}
	return int64(a) + int64(b) + int64(c), nil
}
