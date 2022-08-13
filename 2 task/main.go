package main

import (
	"fmt"
	"os"
)

func main() {
	//если порядок вывода важен, можно использовать waitgroup и массив для хранения результата
	var arr = [5]int{2, 4, 6, 8, 10} //используется фиксированный массив для оптимизации производительности
	ch := make(chan int)             //канал для передачи квадратов каждого из числа
	for i := 0; i < 5; i++ {
		go func(i int) { //запускаем 5 горутин для подсчета квадратов с записью в канал
			ch <- arr[i] * arr[i]
		}(i)
	}
	//этот цикл идет параллельно горутинам и считываем результаты из канала
	for i := 0; i < 5; i++ {
		fmt.Fprintln(os.Stdout, <-ch) // fmt.Println(..) также использует stdout, можно было его не указывать
	}

}