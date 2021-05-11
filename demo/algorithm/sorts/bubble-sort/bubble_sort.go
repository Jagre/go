package main

import(
	"fmt"
)

func main(){

	a:= []int{6,8,7,1,9,2,5,4, 3,10}

	sort(a)

	fmt.Printf("%v", a)
}


/*
时间复杂度：O(n^2)
优点：
1.简单

缺点：
1. 慢
2. 不稳定（相同数字的顺序不确定）
*/
func sort(a []int){
	var length = len(a);

	for i:=0; i < length -1 ; i++{
		for j:=i+1 ; j< length ; j++ {
			if a[i] > a[j]{
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}