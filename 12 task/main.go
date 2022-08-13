package main

import "fmt"

//один из способов - сделать слайс и добавлять туда только уникальные значения
//другой способ - создать map, где ключом является нужный нам тип, а значение - struct{}{}
type MySet struct { //первый способ - структура содержит слайс строк
	Slice []string //для правильной работы со множеством нужно дописать функцию добавления
}

//для добавления новой строки в такую структуру-множество нужна проверка уникальности
func IsUnique(s string, arr []string) bool {
	for _, x := range arr { //обходим все строки массива и каждую сравниваем с новой
		if x == s { //если хоть одна совпала
			return false //не добавим такую строчку
		}
	}
	return true //иначе эта строчка является уникальной и ее можно добавить во множество
}

//сама функция добавления во множество
func AddItem(set *MySet, s string) {
	if IsUnique(s, set.Slice) { //если значение уникально
		set.Slice = append(set.Slice, s) //добавляем его в слайс
	}
}
func main() {
	strs := []string{"cat", "cat", "dog", "cat", "tree"}
	var set MySet //первый способ со структурой-множеством
	for _, x := range strs {
		AddItem(&set, x)
	}
	fmt.Println("Структура-множество: ")
	fmt.Println(set.Slice)
	var mapset = make(map[string]struct{}) //второй вариант
	for _, x := range strs {
		mapset[x] = struct{}{}
	}
	fmt.Println("Карта-множество: ")
	fmt.Println(mapset) //вывод выглядит немного более громоздким, но реализация проще
}
