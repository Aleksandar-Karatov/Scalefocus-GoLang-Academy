package main

import (
	"fmt"
)

func main() {
	inputs := []int{1, 17, 34, 56, 2, 8}
	evenCh := processEven(inputs)
	oddCh := processOdd(inputs)
	close(evenCh)
	close(oddCh)

	fmt.Println("Even:")
	for num := range evenCh {
		fmt.Println(num)
	}

	fmt.Println("Odd:")
	for num := range oddCh {
		fmt.Println(num)
	}
}

func processEven(inputs []int) chan int {
	ch := make(chan int, len(inputs))
	for _, item := range inputs {
		go func(item int) {
			if item%2 == 0 {
				ch <- item
			}
		}(item)
	}
	return ch
}

func processOdd(inputs []int) chan int {
	ch := make(chan int, len(inputs))
	for _, item := range inputs {
		go func(item int) {
			if item%2 != 0 {
				ch <- item
			}
		}(item)
	}
	return ch
}
