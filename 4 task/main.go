package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func Worker(id int, datachan chan int, group *sync.WaitGroup) {
	for { //в бесконечном цикле получаем данные и выводим на экран
		data := <-datachan
		fmt.Printf("ID: %d, data: %d\n", id, data)
	}
}

func main() {
	var group sync.WaitGroup             //для работы горутин как единой группы
	datachan := make(chan int)           //канал с данными
	exitSig := make(chan os.Signal)      //в данный канал запишется событие нажатия ctrl+c
	signal.Notify(exitSig, os.Interrupt) //ctrl+c это зарезервированный сигнал прерывания
	go func() {
		<-exitSig //считываем сигнал
		fmt.Println("\nВыход...")
		os.Exit(0) //завершаем программу
	}()
	var N int
	fmt.Println("Введите число воркеров: ")
	fmt.Fscan(os.Stdin, &N)
	fmt.Println("Для завершения программы нажмите CTRL+C")
	for i := 0; i < N; i++ {
		group.Add(1)
		go Worker(i, datachan, &group) //это горутины для чтения данных с канала
	}
	go func() { //а это горутина для генерации данных  в канал
		for i := 0; ; i++ {
			datachan <- i
			time.Sleep(time.Second * 1) //перерыв на секунду
		}
	}()
	group.Wait()
}
