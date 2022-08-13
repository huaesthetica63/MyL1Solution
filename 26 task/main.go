package main

import (
	"fmt"
	"strings"
)

//первый способ
func CheckWord1(s string) bool {
	arr := []rune(strings.ToUpper(s))
	//получили массив символов, приведенных к верхнему регистру
	//обходим символы и проверяем их на уникальность
	for i, x := range arr {
		for j := i + 1; j < len(arr); j++ {
			if x == arr[j] {
				return false //если есть совпадение по буквам
			}
		}
	}
	return true //если все буквы оказались уникальными
}

//второй способ
func CheckWord2(s string) bool {
	arr := []rune(strings.ToUpper(s))
	//получили массив символов, приведенных к верхнему регистру
	//здесь начались отличия от первого способа - теперь считаем число вхождений
	for _, x := range arr {
		charcount := strings.Count(string(arr), string(x))
		if charcount > 1 {
			return false
		}
	}
	return true //если все буквы оказались уникальными
}
func main() {
	var (
		ex1 = "abcd"
		ex2 = "abCdefAaf"
		ex3 = "aQbcd"
	)
	fmt.Println("Первый способ:")
	fmt.Printf("Строка: %s - %t\n", ex1, CheckWord1(ex1))
	fmt.Printf("Строка: %s - %t\n", ex2, CheckWord1(ex2))
	fmt.Printf("Строка: %s - %t\n", ex3, CheckWord1(ex3))
	fmt.Println("Второй способ:")
	fmt.Printf("Строка: %s - %t\n", ex1, CheckWord2(ex1))
	fmt.Printf("Строка: %s - %t\n", ex2, CheckWord2(ex2))
	fmt.Printf("Строка: %s - %t\n", ex3, CheckWord2(ex3))
}
