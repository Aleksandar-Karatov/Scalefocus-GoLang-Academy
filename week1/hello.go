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

}
