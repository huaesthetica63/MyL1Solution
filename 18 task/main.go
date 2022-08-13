package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//структура-счетчик, внутри содержит беззнаковое целое
type MyCounter struct {
	Count uint64
}

//используем атомарный счетик (пакет atomic)
//gous - число горутин, incrmax - до какого значения каждая из них считает
func (c MyCounter) StartCounting1(gous, incrMax int) {
	var group sync.WaitGroup //заставляем все горутины работать как единую группу
	for i := 0; i < gous; i++ {
		group.Add(1)
		go func() {
			for j := 0; j < incrMax; j++ {
				atomic.AddUint64(&c.Count, 1) //безопасный инкремент в конкурентной среде
				//атомарный счетчик позволяет обезопасить от попытки нескольких рутин обратиться к переменной
			}
			group.Done() //завершаем горутину
		}()

	}
	group.Wait()                                   //ждем пока все горутины закончат работу
	fmt.Printf("Значение счетчика: %d\n", c.Count) //печатаем значение переменной
}

//метод с использованием мьютексов
func (c MyCounter) StartCounting2(gous, incrMax int) {
	var group sync.WaitGroup //заставляем все горутины работать как единую группу
	var mutex sync.Mutex     //мьютекс для блокировки доступа к переменной на время работы горутины
	for i := 0; i < gous; i++ {
		group.Add(1)
		go func() {
			for j := 0; j < incrMax; j++ {
				mutex.Lock() //блокируем доступ другим горутинам
				c.Count++
				mutex.Unlock() //теперь к переменной снова можно обращаться
			}
			group.Done() //завершаем горутину
		}()

	}
	group.Wait()                                   //ждем пока все горутины закончат работу
	fmt.Printf("Значение счетчика: %d\n", c.Count) //печатаем значение переменной
}
func main() {
	var counter MyCounter
	fmt.Printf("Конкурентная среда из %d горутин по %d итераций\n", 500, 100000)
	counter.StartCounting1(500, 100000)
	fmt.Printf("Конкурентная среда из %d горутин по %d итераций\n", 200, 555)
	counter.StartCounting2(200, 555)
}
