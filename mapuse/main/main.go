package main

import (
	"fmt"
)

func main() {
	//第一种使用方式
	var a map[string]string
	//在使用map前，需要先make，make的作用就是给map分配数据空间
	a = make(map[string]string, 10)
	a["no1"] = "宋江" //ok?
	a["no2"] = "吴用"
	a["no1"] = "武松"
	a["no3"] = "吴用"
	fmt.Println(a)

	//第二种使用方式
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)

	//第三种方式
	heroes := map[string]string{
		"hero1": "宋江",
		"hero2": "卢俊义",
		"hero3": "吴用",
	}
	fmt.Println("heroes = ", heroes)

	//案例
	/**
	演示一个key-value 的value是map的案例
	*/
	studentMap := make(map[string]map[string]string)
	studentMap["stud01"] = make(map[string]string, 3)
	studentMap["stud01"]["name"] = "tom"
	studentMap["stud01"]["sex"] = "male"
	studentMap["stud01"]["address"] = "北京长安街~"

	fmt.Println(studentMap)
	fmt.Println(studentMap["stud01"])
}
