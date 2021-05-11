package main

import "fmt"

func main() {
	a := []int{6, 8, 7, 1, 9, 2, 5, 4, 3, 10}

	sort(a)

	fmt.Printf("%v", a)
}

func sort(a []int) {

	length := len(a)
	h := 1
	for h < length/3+1 {
		h *= 2
	}

	for ; h > 0; h /= 2 {
		for i := h; i < length; i += h {
			for j := i; j > 0; j -= h {
				if a[j] < a[j-h] {
					a[j], a[j-h] = a[j-h], a[j]
				} else {
					break
				}
			}
		}
	}
}
