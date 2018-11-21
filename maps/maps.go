package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

/*
map将键映射到值。
map的零值为 nil 。nil 映射既没有键，也不能添加键。
make 函数会返回给定类型的map，并将其初始化备用。
*/
var m map[string]Vertex

func main() {

	fmt.Println(m == nil) // true

	m = make(map[string]Vertex)
	fmt.Println(m == nil) // false

	m["Bell Labs"] = Vertex{
		40.68443, -74.39967,
	}
	fmt.Println(m["Bell Labs"]) // {40.68443 -74.39967}

	// 映射的文法与结构体相似，不过必须有键名。
	var map1 = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68443, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(map1) // map[Bell Labs:{40.68443 -74.39967} Google:{37.42202 -122.08408}]

	// 若顶级类型只是一个类型名，你可以在文法的元素中省略它。
	// 下面省略了Vertex类型
	var map2 = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(map2) // map[Bell Labs:{40.68443 -74.39967} Google:{37.42202 -122.08408}]

}
