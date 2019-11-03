package starting

import (
	"fmt"
	"sort"
)

func SortSlice(a, b, c, d int) {
	sl := make([]int, 4, 4)
	sl[0] = a
	sl[1] = b
	sl[2] = c
	sl[3] = d

	sort.Ints(sl)

	for i := range sl {
		fmt.Println(sl[i])
	}
}
