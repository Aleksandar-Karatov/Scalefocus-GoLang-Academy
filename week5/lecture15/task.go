package main

import (
	"fmt"
	"log"
	"sort"
	"time"
)

func main() {
	dates := []string{"Sep-14-2008", "Dec-03-2021", "Mar-18-2022"}
	orderedDates, err := sortDates("Jan-02-2006", dates...)
	if err != nil {
		log.Fatalln(err)
	} else {
		for _, date := range orderedDates {
			fmt.Println(date)

		}
	}

	datesWithErr := []string{"Sep-14-2008", "Dec-03-2021", "Mar-18-2022", "err"}
	orderedDates, err = sortDates("Jan-02-2006", datesWithErr...)
	if err != nil {
		log.Fatalln(err)
	} else {
		for _, date := range orderedDates {
			fmt.Println(date)

		}
	}
}
func sortDates(format string, dates ...string) ([]string, error) { /* implement */
	var datesInTime []time.Time
	var datesInOrder []string
	for _, date := range dates {

		time, err := time.Parse(format, date)
		if err != nil {
			return nil, err
		} else {
			datesInTime = append(datesInTime, time)
		}
	}
	sort.Slice(datesInTime, func(i, j int) bool {
		return datesInTime[i].Before(datesInTime[j])
	})
	for _, date := range datesInTime {
		datesInOrder = append(datesInOrder, date.Format(format))
	}

	return datesInOrder, nil

}
