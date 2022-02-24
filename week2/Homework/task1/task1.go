package main

import "fmt"

func main() {
	fmt.Println(daysInMonth(2, 2020))
	fmt.Println(daysInMonth(2, 2015))
	fmt.Println(daysInMonth(2, 2016))
	fmt.Println(daysInMonth(2, 300))
	fmt.Println(daysInMonth(1, 2020))
	fmt.Println(daysInMonth(3, 2020))
	fmt.Println(daysInMonth(4, 2020))
	fmt.Println(daysInMonth(5, 2020))
	fmt.Println(daysInMonth(6, 2020))
	fmt.Println(daysInMonth(7, 2020))
	fmt.Println(daysInMonth(8, 2020))
	fmt.Println(daysInMonth(9, 2020))
	fmt.Println(daysInMonth(10, 2020))
	fmt.Println(daysInMonth(11, 2020))
	fmt.Println(daysInMonth(12, 2020))
	fmt.Println(daysInMonth(14, 2020))
	fmt.Println(daysInMonth(-11, 2020))
	fmt.Println(daysInMonth(0, 2020))

}

func daysInMonth(month int, year int) (int, bool) {

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31, true
	case 4, 6, 9, 11:
		return 30, true
	case 2:
		return (howManyDaysInFebruary(year)), true
	default:
		return -1, false
	}

}

func howManyDaysInFebruary(year int) int {
	if (year%4 == 0 && year%100 == 0 && year%400 == 0) || (year%4 == 0 && year%100 != 0) {
		return 29
	} else {
		return 28
	}
}
