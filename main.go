package main

import (
	"coursera_golang/fmi"
	"fmt"
)

func main() {
	arr := []int{10, 20, 51, 21, 1, 231, 314, 124, 34, 21, 519, 124, 653, 121, 4612}

	fmi.BubbleSort(arr)

	fmt.Println(arr)
}
