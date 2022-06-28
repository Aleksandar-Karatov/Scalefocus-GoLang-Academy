package main

import (
	"log"
	"time"
)

func main() {
	out := generateThrottled("foo", 2, time.Second)
	for f := range out {
		log.Println(f)
	}
}
func generateThrottled(data string, bufferLimit int, clearInterval time.Duration) <-chan string {

	ch := make(chan string, bufferLimit)

	temp := bufferLimit
	go func() {
		for {
			if temp > 0 {
				ch <- data
				temp--
			} else {
				time.Sleep(clearInterval)
				temp = bufferLimit
			}

		}
	}()

	return ch

}
