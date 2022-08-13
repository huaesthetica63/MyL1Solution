package main

import (
	"context"
	"fmt"
	"time"
)

//первый способ основан на использовании канала, который является сигналом
//как только в этот канал помещается сигнал (булева переменная), то горутина останавливается
func FirstWay(seconds int) { //аргумент - число секунд до отправки стоп-сигнала
	finish := make(chan bool) //маркер завершения горутины (небуф. канал)
	go func() {               //сама горутина
		for { //бесконечный цикл "прослушки канала"
			select {
			case <-finish: //если в канал что-то положили
				fmt.Println("Стоп...")
				return //выходим из горутины
			default: //во всех других случаях
				fmt.Println("Горутина 1 что-то делает...")
				time.Sleep(time.Second) //чтобы не забивал консоль очень часто
			}
		}
	}()
	fmt.Printf("Горутина 1 остановится через %d сек.\n", seconds)
	time.Sleep(time.Second * time.Duration(seconds)) //даем горутине поработать
	finish <- true                                   //сигнал о завершении
	fmt.Println("Горутина 1 остановилась!")
}

//второй способ тоже основан на использовании канала
//на этот раз канал нужно закрыть, чтобы остановилась горутина
func SecondWay() {
	finish := make(chan string) //канал, который в нужный момент должен закрыться
	defer fmt.Println("Горутина 2 остановилась!")
	go func() {
		for { //бесконечный цикл
			res, ok := <-finish //пытаемся обратиться к каналу
			if ok {             //если получить данные удалось
				fmt.Println("Получено сообщение: " + res)
			} else { //если получили ошибку, значит пора завершаться
				fmt.Println("Стоп...")
				return
			}
		}
	}()
	fmt.Println("Горутина 2 стартовала")
	finish <- "Еще не финиш..."
	time.Sleep(time.Second)
	close(finish) //закрываем канал
}

//третий способ основан на использовании контекста, который даст "сигнал" о завершении горутины
func ThirdWay(seconds int) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	fin := make(chan bool) //чтобы основная функция подождала завершение горутины
	go func(fin chan bool) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Стоп...")
				fin <- true
				return
			default:
				fmt.Println("Горутина 3 что-то делает...")
				time.Sleep(time.Second)
			}
		}
	}(fin)
	fmt.Printf("Горутина 3 остановится через %d сек.\n", seconds)
	time.Sleep(time.Second * time.Duration(seconds)) //даем горутине поработать
	cancelFunc()                                     //сигнал о завершении
	<-fin                                            //теперь можем закрыть функцию
	fmt.Println("Горутина 3 остановилась!")
}

func main() {
	FirstWay(5) //метод остановки через канал-сигнал
	SecondWay() // метод остановки через закрытие канала
	ThirdWay(5) //метод с контекстами
}
