package main

import (
	"fmt"
	"sync"
	"time"
)

type response struct {
	Msg string
	Err error
}

func craw(a chan response, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	a <- response{Msg: fmt.Sprint("time ", time.Now().UnixMilli()), Err: nil}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	a := make(chan response, 10)

	for i := 0; i < 10; i++ {
		go craw(a, &wg)
	}

	wg.Wait()

	for i := 0; i < 10; i++ {
		fmt.Println(<-a)
	}

}
