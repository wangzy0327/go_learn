package main

import (
	"fmt"
	"sort"
)

func main() {
	var map1 map[int]string
	map1 = make(map[int]string)
	map1[1] = "东方不败"
	map1[2] = "风清扬"
	map1[3] = "方证大师"
	map1[4] = "令狐冲"

	var keys []int
	for k := range map1 {
		keys = append(keys, k)
	}
	//排序
	sort.Ints(keys)
	fmt.Println(keys)

	for _, k := range keys {
		fmt.Printf("map[%v] = %v\n", k, map1[k])
	}

}
