package main

import "fmt"

//из исходного слайса получаем канал на чтение из него данных (односторонний)
func GenerateChan(arr []int) <-chan int {
	newchan := make(chan int) //новый канал
	go func() {               //запускаем горутину, которая будет заполнять канал числами массива
		for _, x := range arr {
			newchan <- x //записываем в канал число
		}
		close(newchan) //закрываем канал после того, как вторая часть конвеера считает оттуда значения
	}()
	return newchan // возвращаем канал
}

//вторая часть конвеера, которая считает квадраты чисел из канала
func CountChan(numchan <-chan int) <-chan int {
	reschan := make(chan int) //канал с квадратами чисел
	go func() {               //горутина для подсчета квадратов
		for x := range numchan {
			reschan <- x * x //записываем квадрат в новый канал
		}
		close(reschan) //в main все отсюда считаем и закроем канал
	}()
	return reschan
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1213}
	fmt.Println("Исходный массив:")
	fmt.Println(arr)
	chan1 := GenerateChan(arr) //запускаем горутину по чтению исходных чисел
	chan2 := CountChan(chan1)  //и горутину по вычислению квдаратов
	fmt.Println("Квадраты чисел массива:")
	for n := range chan2 {
		fmt.Println(n) //выписываем квадраты на экран
	}
}
