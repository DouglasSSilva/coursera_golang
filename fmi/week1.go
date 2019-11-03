package fmi

func BubbleSort(arr []int) {

	arrLen := len(arr)
	for i := range arr {
		for k := i + 1; k < arrLen; k++ {
			if arr[i] > arr[k] {
				arr[i], arr[k] = arr[k], arr[i]
			}
		}
	}
}
