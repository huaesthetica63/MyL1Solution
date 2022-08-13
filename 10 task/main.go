package main

import "fmt"

func main() {
	arr := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println("Исходный массив:")
	fmt.Println(arr)
	//первый способ: делим число на 10 и приводим к целому числу
	//так, числа (-10;10) попадут в 0, [10,20) в 10, [20,30) в 20, (-30,20] в -20
	MyMap1 := make(map[int][]float32)
	for _, x := range arr {
		intApprox := int(x) / 10
		MyMap1[intApprox*10] = append(MyMap1[intApprox*10], x)
	}
	fmt.Println("Первый способ разделения чисел:")
	fmt.Println(MyMap1)
	//второй способ: -30: числа больше -30 но до -20,
	//-20: числа больше -20, но меньше -10, -10: до 0, 0: от 0 до 10 ....
	MyMap2 := make(map[int][]float32)
	for _, x := range arr {
		intApprox := int(x) / 10
		if x < 0 { //здесь отличия от первого способа в разделении отрицательных чисел
			intApprox--
		}
		MyMap2[intApprox*10] = append(MyMap2[intApprox*10], x)
	}
	fmt.Println("Второй способ разделения чисел:")
	fmt.Println(MyMap2)
}
