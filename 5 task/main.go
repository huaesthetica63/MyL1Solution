package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func Worker(inputchan chan int, ctx context.Context) { //обработка данных из канала
	for { //бесконечный цикл
		select {
		case <-ctx.Done(): //если истек таймер
			{
				fmt.Println("Таймер истек!")
				os.Exit(0) //выходим из всей программы
			}
		case data := <-inputchan: //считываем данные с канала
			fmt.Println(data)
		}
	}
}
func main() {
	var N int
	fmt.Println("Введите N: ")
	fmt.Fscan(os.Stdin, &N)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(N)) //контекст завершится по истечению таймера
	defer cancel()                                                                         //контекст завершится в любом случае (этой командой предостерегаем программу от "незавершения" по таймеру)
	datachan := make(chan int)                                                             //канал обмена данными
	go Worker(datachan, ctx)                                                               //запускаем горутину
	for i := 0; ; i++ {                                                                    //заполняем канал последовательностью чисел, добавляя новое каждую секунду
		datachan <- i
		time.Sleep(time.Second)
	}
}
