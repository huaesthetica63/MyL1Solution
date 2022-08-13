package main

import (
	"fmt"
)

var justString string

//создает строку заданной длины

//строка - набор произвольных байтов
//для того, чтобы создать строку с заданным количеством символов,
//необходимо создать слайс рун заданной длины
func createHugeString(size int64) string {
	arr := make([]rune, size)
	for i, _ := range arr {
		arr[i] = 'ä' //юникод-символ
	}
	return string(arr)
}
func someFunc() {
	v := createHugeString(1 << 10)     //создаем строку с заданным числом символов
	runeArr := []rune(v)               //преобразуем string в руны
	justString = string(runeArr[:100]) //выделяем 100 символов из массива рун
	//если мы оставим код из примера, символы длиннее одного байта (юникод) выделятся не полностью
	//то есть, мы получим не 100 символов, а 50 (100 байтов по 2)
}

func main() {
	fmt.Println("Вызов функции someFunc...")
	someFunc()
	fmt.Println("Полученная строка: ")
	fmt.Println(justString)
}
