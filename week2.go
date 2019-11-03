package main

import (
	"fmt"
	"strconv"
	"strings"
)

func truncate(val string) {

	floated, _ := strconv.ParseFloat(val, 64)

	fmt.Printf("%d\n", int(floated))

}

func findian(val string) {
	i := "i"
	a := "a"
	n := "n"

	if string(val[0]) == i && string(val[len(val)-1]) == n && strings.Contains(val, a) {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
	}
}
