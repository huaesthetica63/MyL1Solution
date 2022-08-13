package main

import (
	"fmt"
	"sync"
)

func main() {
	mymap := make(map[int]int)  //создаем карту, ключ и значение - число
	var group sync.WaitGroup    // группа горутин, чтобы main не завершился без них
	var mutex sync.Mutex        // мьютекс для блокировки доступа к записи
	gous := 500                 // число горутин
	for i := 1; i < gous; i++ { //запускаем каждую из них
		group.Add(1)
		go func(i int) {
			mutex.Lock()
			mymap[i] = i
			mutex.Unlock()
			group.Done()
		}(i)
	}
	group.Wait()
	fmt.Println(mymap)
}
