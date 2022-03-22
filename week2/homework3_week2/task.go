package main

import "fmt"

type MagicList struct {
	LastItem *Item
}
type Item struct {
	Value    int
	PrevItem *Item
}

func main() {
	l := &MagicList{} // tests from the presentation plus a few bonus tests
	add(l, 10)
	add(l, 22)
	add(l, 44)
	add(l, 55)
	add(l, 66)
	add(l, 77)
	add(l, 88)
	add(l, 99)
	add(l, 111)
	add(l, 222)
	add(l, 333)
	add(l, 444)
	fmt.Println(toSlice(l))
}
func add(l *MagicList, value int) { // function to add items to an existing MagicList
	var newItem = Item{value, nil}
	if l.LastItem == nil {
		l.LastItem = &newItem
	} else {
		newItem.PrevItem = l.LastItem
		l.LastItem = &newItem
	}

}
func toSlice(l *MagicList) []int {
	var slicedList []int
	fillSlice(l.LastItem, &slicedList) // I was cheeky and used a helper function so that I can use recursion
	for i, j := 0, len(slicedList)-1; i < j; i, j = i+1, j-1 {
		slicedList[i], slicedList[j] = slicedList[j], slicedList[i]
	}
	return slicedList

}
func fillSlice(item *Item, slicedItems *[]int) { // recursive function to fill slice with item values
	if item == nil {
		return
	} else {
		*slicedItems = append(*slicedItems, item.Value)
		fillSlice(item.PrevItem, slicedItems)
	}
}
