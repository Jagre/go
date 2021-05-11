package main

import "fmt"

func main(){
	a:= []int{6,8,7,1,9,2,5,4, 3,10}

	sort(a);

	fmt.Printf("%v", a)
}

func sort(a []int){
	length := len(a)
	if length < 2{
		return
	}

	for i:= 1; i < length; i++{
		for j:= i; j > 0; j--{
			if a[j] < a[j-1]{
				a[j], a[j-1] = a[j-1], a[j]
			}else {
				break
			}
		}
	}
}