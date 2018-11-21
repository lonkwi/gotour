package main

import "fmt"

/*
for 循环的 range 形式可遍历切片或映射。
当使用 for 循环遍历切片时，每次迭代都会返回两个值。
第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。
*/
func main1() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

/*
可以将下标或值赋予 _ 来忽略它。
若你只需要索引，去掉 , value 的部分即可。
注：索引和value都可以省略
*/
func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // =2**i
	}

	// 省略索引
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	// 省略value
	for i, _ := range pow {
		fmt.Printf("%d\n", i)
	}
}
