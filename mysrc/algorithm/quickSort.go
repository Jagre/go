package main

import (
	"fmt"
)

func main() {
	arr := [6]int{6, 2, 7, 3, 8, 9}
	quickSort(arr[:], 0, len(arr)-1)

	fmt.Print(arr)
}

func quickSort(arr []int, left, right int) {
	i, j := left, right
	if i >= j {
		return
	}
	key := arr[i]
	for i < j {
		for i < j && key <= arr[j] {
			j--
		}
		arr[i] = arr[j]

		for i < j && key >= arr[i] {
			i++
		}
		arr[j] = arr[i]
	}
	arr[i] = key //****重点，最精妙：一轮遍历后，key归位
	//这时索引i的值，已确定其位置顺序（左边都比i小，右边都比i大）
	quickSort(arr, left, i-1)
	quickSort(arr, i+1, right)
}
