package main

import "fmt"

//первый метод удаления через append
func DeleteElement1(slice []int, ind int) []int {
	if ind > len(slice) || ind <= 0 { //проверяем корректность индекса
		panic("Нет такого элемента!")
	}
	res := append(slice[:ind-1], slice[ind:]...)
	//мы не включили в res как раз slice[ind], многоточие - множественный аргумент
	//в функции append (на самом деле вместо слайса "подставляются" все его элементы по отдельности)
	return res
} //в первом методе сохранился порядок следования элементов слайса
//второй метод не сохраняет порядок
func DeleteElement2(slice []int, ind int) []int {
	if ind > len(slice) || ind <= 0 { //проверяем корректность индекса
		panic("Нет такого элемента!")
	}
	slice[ind-1] = slice[len(slice)-1] //элементу, который хотим удалить, присваиваем конец слайса
	return slice[:len(slice)-1]        //возвращаем срез без конечного элемента (отрезали один элемент)
}
func main() {
	exampleSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Изначальный слайс: ")
	fmt.Println(exampleSlice)
	exampleSlice = DeleteElement1(exampleSlice, 4) //удалили 4 элемент
	fmt.Println("Удалили 4-й элемент, метод 1:")
	fmt.Println(exampleSlice)
	exampleSlice = DeleteElement2(exampleSlice, 8) //удалили 8 элемент
	fmt.Println("Удалили 8-й элемент, метод 2:")
	fmt.Println(exampleSlice)
}
