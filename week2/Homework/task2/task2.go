package main

import (
	"fmt"
	"math/rand"
	"time"
)

func citiesAndPrices() ([]string, []int) {
	rand.Seed(time.Now().UnixMilli())
	cityChoices := []string{"Berlin", "Moscow", "Chicago", "Tokyo", "London"}
	dataPointCount := 100

	// randomly choose cities
	cities := make([]string, dataPointCount)
	for i := range cities {
		cities[i] = cityChoices[rand.Intn(len(cityChoices))]
	}

	prices := make([]int, dataPointCount)
	for i := range prices {
		prices[i] = rand.Intn(100)
	}

	return cities, prices
}
func main() {
	cities, prices := citiesAndPrices()

	pricedCities := groupSlices(cities, prices)
	for city := range pricedCities {
		fmt.Println(city, pricedCities[city])
	}
	//fmt.Println(pricedCities)
}

func groupSlices(keySlice []string, valueSlice []int) map[string][]int {
	var citiesMappedToPrices = make(map[string][]int)
	for i := 0; i < len(keySlice); i++ {
		citiesMappedToPrices[keySlice[i]] = append(citiesMappedToPrices[keySlice[i]], valueSlice[i])
	}

	return citiesMappedToPrices
}
