package main

import "fmt"

//возвращаем индекс найденного числа или -1, если ничего не удалось найти
func BinarySearch(arr []int, searchVal int) int {
	var (
		left  = 0            //крайняя левая граница - начало слайса
		right = len(arr) - 1 //крайняя правая - конец
	)
	//на каждой итерации этот промежуток от left до right будет половиниться
	for left <= right {
		m := (left + right) / 2 //середина промежутка
		if arr[m] < searchVal { //если середина меньше требуемого значения
			left = m + 1 //половиним промежуток в правую сторону
		} else {
			right = m - 1 //иначе - в левую
		}
	}
	if left >= len(arr) {
		return -1 //если вышли за пределы слайса и ничего не нашли
	}
	if arr[left] != searchVal {
		return -1 //если оставшееся значение не совпало с требуемым
	}
	return left //иначе left-индекс найденного элемента
}
func main() {
	arr := []int{-12, 0, 11, 24, 66, 2232} //изначальный слайс должен быть отсортирован
	fmt.Println("Исходный массив: ")
	fmt.Println(arr)
	search := 66
	ind := BinarySearch(arr, search)
	fmt.Printf("Число 66 - элемент массива по номеру: %d\n", ind+1)
}
