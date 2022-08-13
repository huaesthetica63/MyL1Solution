package main

import (
	"fmt"
	"math"
)

//инкапсуляция полей-координат отражена в их написании с маленькой буквы
type Point struct {
	x float64
	y float64
}

//конструктор с параметрами
func NewPoint(a, b float64) *Point {
	res := new(Point)
	res.x = a
	res.y = b
	return res
}

//геттеры для получения инкапсулированных координат
func (a Point) GetX() float64 {
	return a.x
}
func (a Point) GetY() float64 {
	return a.y
}

//поиск расстояния по формуле Пифагора
func FindS(a, b Point) float64 {
	x := math.Sqrt(math.Pow(a.GetX()-b.GetX(), 2) + math.Pow(a.GetY()-b.GetY(), 2))
	return x
}

//функция наглядного вывода координат
func (p Point) Print() {
	//.2f - выводим только 2 точки после запятой для аккуратного вывода на экран
	fmt.Printf("Точка: (%.2f; %.2f)\n", p.GetX(), p.GetY())
}
func main() {
	a := *NewPoint(1, 5)
	b := *NewPoint(4, 9)
	a.Print()
	b.Print()
	fmt.Printf("Расстояние между точками: %.2f\n", FindS(a, b))
}
