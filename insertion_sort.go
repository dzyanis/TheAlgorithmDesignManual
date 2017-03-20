package main

import (
	"fmt"
)

func insertionSort(list []byte) []byte {
	for i := 0; i<len(list); i++ {
		for j := i; j>0 && list[j] < list[j-1]; j-- {
			list[j], list[j-1] = list[j-1], list[j]
		}
	}
	return list
}

func main() {
	list := []byte("INSERTIONSORT")
	fmt.Printf("In:  %s\n", list)
	newList := insertionSort(list)
	fmt.Printf("Out: %s\n", string(newList))

	list = []byte("37495261")
	fmt.Printf("In:  %s\n", list)
	newList = insertionSort(list)
	fmt.Printf("Out: %s\n", string(newList))
}
