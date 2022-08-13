package main

import "fmt"

//множество представляет собой структуру со слайсом, добавляются только уникальных значений
type MySet struct {
	Slice []string
}

func IsUnique(s string, arr []string) bool {
	for _, x := range arr { //обходим весь срез
		if x == s { //проверяем каждый элемент на совпадение с добавляемой строкой
			return false //если совпало хоть раз - такой элемент уже есть и не будет добавлен снова
		}
	}
	return true
}

//сама функция добавления во множество
func AddItem(set *MySet, s string) {
	if IsUnique(s, set.Slice) { //если значение уникально
		set.Slice = append(set.Slice, s) //добавляем его в слайс
	}
}

//функция получения нового множества, содержащего элементы и из a, и из b одновременно
func FindCross(a, b MySet) MySet {
	var res MySet
	//обходим все элементы первого множества, ищем их во втором, если нашли - добавляем в пересечение
	for _, x := range a.Slice {
		if !IsUnique(x, b.Slice) { //если элемент из а попался в b
			res.Slice = append(res.Slice, x) //это элемент пересечения
		}
	}
	return res
}
func main() {
	strs1 := []string{"cat", "cat", "dog", "cat", "tree", "cross"}
	strs2 := []string{"cat", "cross", "set", "get", "string", "cross"}
	//реализация со структурами
	var set1, set2 MySet
	for _, x := range strs1 {
		AddItem(&set1, x)
	}
	for _, x := range strs2 {
		AddItem(&set2, x)
	}
	fmt.Println("Множество а:")
	fmt.Println(set1.Slice)
	fmt.Println("Множество b:")
	fmt.Println(set2.Slice)
	res := FindCross(set1, set2)
	fmt.Println("Пересечение: ")
	fmt.Println(res.Slice)

	//те же множества, но уже в реализации с картами:
	var mapset1 = make(map[string]struct{})
	var mapset2 = make(map[string]struct{})
	for _, x := range strs1 {
		mapset1[x] = struct{}{}
	}
	for _, x := range strs2 {
		mapset2[x] = struct{}{}
	}
	var crossmap = make(map[string]struct{})
	for key := range mapset1 { //обходим все ключи карты
		_, ok := mapset2[key] //ок- нашли или нет запись по ключу во второй карте
		if ok {               //если нашли, этот ключ общий для двух карт
			crossmap[key] = struct{}{} //добавляем его в карту-множество
		}
	}
	fmt.Println("Множество-карта а: ")
	fmt.Println(mapset1)
	fmt.Println("Множество-карта b: ")
	fmt.Println(mapset2)
	fmt.Println("Пересечение карт:")
	fmt.Println(crossmap) //выводим результат в реализации карт
}
