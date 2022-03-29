package main

import (
	"fmt"
	"sync"
	"time"
)

type ConcurrentPrinter struct {
	sync.WaitGroup
	sync.Mutex /* Feel free to add your own state here */
}

func (cp *ConcurrentPrinter) PrintFoo(times int) {
	cp.WaitGroup.Add(times / 2)
	go func() {

		for i := 0; i < times/2; i++ {
			time.Sleep(10 * time.Millisecond)

			cp.Mutex.Lock()

			fmt.Println("foo")

			cp.WaitGroup.Done()
			cp.Mutex.Unlock()
			time.Sleep(10 * time.Millisecond)
			//time.Sleep(1 * time.Microsecond)

		}
	}()
}
func (cp *ConcurrentPrinter) PrintBar(times int) {
	cp.WaitGroup.Add(times / 2)

	go func() {
		for i := 0; i < times/2; i++ {
			time.Sleep(10 * time.Millisecond)

			cp.Mutex.Lock()

			fmt.Println("bar")
			cp.WaitGroup.Done()
			cp.Mutex.Unlock()
			time.Sleep(10 * time.Millisecond)

		}
	}()
}

func main() {
	times := 10
	cp := &ConcurrentPrinter{}
	cp.PrintFoo(times)
	cp.PrintBar(times)
	cp.Wait()
}
