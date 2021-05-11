package main

import (
	"fmt"
)

/*
大数阶乘
问题点：结果大了就会溢出
办法：将结果存入数组，个位放在result[0], 十位放在result[1] , 依此类推
原理：满10就向高位（eg: 个位向十位，result[0]向result[1]）进位，其实细想就跟数学乘法计算列坚式的原理很像
注：result数组的长度根据公式N!的位数=[log(1)+log(2)+…..log(N)]+1([]表示向上取整)
*/

func main() {
	var res [10]int
	res[0] = 1
	carry := 0 //进位
	for i := 2; i < 10; i++ {
		for j := 0; j < len(res); j++ {
			//核心
			res[j] = res[j]*i + carry
			carry = res[j] / 10
			res[j] = res[j] % 10
		}
	}

	for i := len(res) - 1; i >= 0; i-- {
		fmt.Printf("%d", res[i])
	}
}
