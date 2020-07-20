package main

import (
	"fmt"
	"sync"
	"time"
)

func deposit(balance *int, amount int, mtx *sync.Mutex, wg *sync.WaitGroup) {
	mtx.Lock()
	*balance += amount
	mtx.Unlock()
	wg.Done()

}

func withdraw(balance *int, amount int, mtx *sync.Mutex, wg *sync.WaitGroup) {
	mtx.Lock()
	*balance -= amount
	mtx.Unlock()
	wg.Done()
}

func main() {

	balance := 100
	var wg sync.WaitGroup
	var mtx sync.Mutex

	wg.Add(2)
	go deposit(&balance, 10, &mtx, &wg)
	withdraw(&balance, 50, &mtx, &wg)

	wg.Wait()
	fmt.Println(balance)

	mtx2 := sync.Mutex{}
	mtx2.Lock()
	go func() {
		mtx2.Lock() //cannot acquire another lock after already obtaining a lock
		fmt.Println("In goroutine")
		mtx2.Unlock()
	}()

	fmt.Println("In MAIN")
	mtx2.Unlock()               //this release the lock taken on mtx2
	time.Sleep(time.Second * 1) //sleep added to have goroutine run in the

}
