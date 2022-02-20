package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var a int = 15
	var b = true
	var str1, str2 string = "Hello,", "world!"
	fmt.Println(a + 14)
	if b {
		fmt.Println(str1 + " " + str2)
	}
	for i := 0; i < 5; i++ {
		fmt.Println("loop ", i)
	}
	var c = rand.Intn(101)
	switch c {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println(c)
	}

	str1 = str1 + " " + str2
	fmt.Println(str1)
	fmt.Println(rec(6))

	fmt.Println(recFindNeedleInHaystack("haystack", 'p', 0))
	recFizBuzz(1)
}

func rec(num int) int {
	if num == 1 {
		return 1
	} else {
		return num * (rec(num - 1))
	}
}

func recFindNeedleInHaystack(haystack string, needle rune, index int) int {
	if index == (len(haystack) - 1) {
		return -1
	} else if haystack[index] == byte(needle) {
		return index
	} else {
		return recFindNeedleInHaystack(haystack, needle, (index + 1))
	}
}

func recFizBuzz(n int) {
	if n == 101 {
		return
	}
	fmt.Println()
	if n%3 != 0 && n%5 != 0 {
		fmt.Print(n)
	} else {
		if n%3 == 0 {
			fmt.Print("fizz")
		}
		if n%5 == 0 {
			fmt.Print("Buzz")
		}
	}

	recFizBuzz(n + 1)
}
