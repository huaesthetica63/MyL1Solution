package main

import (
	"context"
	"fmt"
	"time"
)

//первый способ с тайм-аутом
func MySleep1(sec int) {
	<-time.After(time.Duration(sec) * time.Second) //дожидаемся истечения таймаута
}

//второй способ с контекстом
func MySleep2(sec int) {
	//создаем контекст с таймаутом
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*time.Duration(sec))
	defer cancelFunc() //по истечению таймера он должен завершиться
	<-ctx.Done()       //ждем завершения контекста
}
func main() {
	seconds := 5
	fmt.Printf("Sleep1 на %d секунд \n", seconds)
	MySleep1(seconds)
	fmt.Println("Все!")
	fmt.Printf("Sleep2 на %d секунд \n", seconds)
	MySleep2(seconds)
	fmt.Println("Все!")
}
