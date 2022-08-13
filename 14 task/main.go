package main

import (
	"fmt"
	"reflect"
)

func sampleFunc() { //для демонстрации переменной типа "функция"

}
func main() {
	//пустой интерфейс соответствует любому типу данных, поэтому его можно использовать при определении типов переменных
	//interface{}.(type) позволяет получить тип данных
	t := []interface{}{123455, "string", sampleFunc, true, 12.2, 12 + 3i, 'c'}
	for _, x := range t { //обход всего слайса
		//type позволяет определять стандартные типы
		switch t := x.(type) {
		default:
			fmt.Printf("%v - это тип %T\n", x, t) //%T - выводит тип анонимного интерфейса(нашей переменной)
		case int:
			fmt.Printf("%v - это целочисленная переменная\n", x)
		case string:
			fmt.Printf("%v - это строка\n", x)
		case bool:
			fmt.Printf("%v - это логический тип\n", x)
		}
	}
	//другой вариант с reflect
	fmt.Println("Определение типов с reflect:")
	for _, x := range t { //обход всего слайса
		//type позволяет определять стандартные типы
		fmt.Printf("%v - тип %v\n", x, reflect.TypeOf(x))
	}

}
