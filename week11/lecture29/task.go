package main

import (
	"fmt"
	"strconv"
)

func main() {

	results := GroupBy([]Order{
		{Customer: "John", Amount: 1000},
		{Customer: "Sara", Amount: 2000},
		{Customer: "Sara", Amount: 1800},
		{Customer: "John", Amount: 1200},
	}, func(o Order) string { return o.Customer }, func(o Order) string { return strconv.Itoa(o.Amount) })
	fmt.Println(results)
}

type Order struct {
	Customer string
	Amount   int
}

func (o Order) TakeAmmount() int {
	return o.Amount
}
func (o Order) TakeCustomer() string {
	return o.Customer
}
func GroupBy[T any, U comparable](col []T, keyFn func(T) U, keyFN2 func(T) U) map[U][]U {
	result := make(map[U][]U)
	for _, item := range col {

		result[keyFn(item)] = append(result[keyFn(item)], keyFN2(item))

	}
	return result

}

func Map[T, K any](ins []T, mapper func(T) K) []K {
	var outs []K
	for _, v := range ins {
		outs = append(outs, mapper(v))
	}
	return outs
}
