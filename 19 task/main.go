package main

import "fmt"

func ReverseWord1(s string) string {
	chars := []rune(s)         //получаем строку в виде массива символов
	l := len(chars)            //получаем длину массива (количество символов)
	for i := 0; i < l/2; i++ { //обходим половину массива
		chars[i], chars[l-1-i] = chars[l-1-i], chars[i] //"разворачиваем массив"
	}
	return string(chars) //переводим массив обратно в строку
}

//второй метод осуществлен чере посимвольную конкатенацию
func ReverseWord2(s string) string {
	result := ""
	for _, x := range s {
		result = string(x) + result //чтобы получить запись задом-наперед,
		//символ прибавляем спереди исходной строки, а не после нее
	}
	return result
}
func main() {
	var abr = "главрыба 개"
	fmt.Println("Оригинальная строка: " + abr)
	fmt.Println("Первый метод: " + ReverseWord1(abr))
	fmt.Println("Второй метод: " + ReverseWord2(abr))
}
