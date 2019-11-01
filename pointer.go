package main

import "fmt"

func pointersCheck() {
	x := 1
	var y int
	var ip *int

	ip = &x
	y = *ip

	fmt.Println(x, *ip, y)

	ptr := new(int)

	*ptr = 3
	fmt.Println(*ptr, ptr)
}
