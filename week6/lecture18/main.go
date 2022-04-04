package main

import (
	"fmt"
	"time"
)

func main() {

	primeNums := goprimesAndSleep(100, 10*time.Millisecond)
	fmt.Println(primeNums)
}

func goprimesAndSleep(n int, sleep time.Duration) []int {
	res := []int{}
	ch := make(chan int, n*n)
	tempChCap := n * n
	go func() {

		for k := 2; k < n; k++ {
			for i := 2; i < n; i++ {

				if len(ch) < cap(ch) {
					if k%i == 0 {
						if k == i {
							ch <- k
							time.Sleep(sleep)
							tempChCap--
						}
						break
					}
				} else {
					close(ch)
					return

				}

			}

		}
		close(ch)
	}()
	for elem := range ch {
		if len(ch) == cap(ch) {
			close(ch)
			return res
		}

		res = append(res, elem)

	}
	return res
}
