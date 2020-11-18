package main

import (
	"fmt"
)

//Person ...
type Person struct {
	Name string
	Age  int
	Sex string
}

func main() {
	//方式1
	var p1 Person
	p1.Age = 20
	p1.Name = "Bob"
	fmt.Println(p1)
	//方式2
	p2 := Person{"mary", 20,"male"}
	//p2.Name = "tom"
	//p2.Age = 18
	fmt.Println(p2)
	//方式3
	var p3 *Person = new(Person)
	//因为p3是一个指针，因此标准的给字段赋值方式
	//(*p3).Name = "smith"
	//原因：go的设计者  为了程序员使用方便，底层会对 p3.Name = "smith"进行处理
	//会给p3 加上取值运算 (*p3).Name = "smith"
	p3.Name = "smith"
	p3.Age = 100
	fmt.Println(*p3)
	//方式4
	//案例：var person *Person = &Person{}
	var person *Person = &Person{}
	//因为person是一个指针，因此标准的访问字段的方法
	//(*person).Name = "scott"
	//go的设计者为了程序员使用方便，也可以person.Name = "scott"
	//原因和上面一样，底层会对 person.Name="scott"进行处理，会加上(*person)
	(*person).Name = "scott"
	person.Name = "scott~"
	person.Age = 50
	fmt.Println(person)
}
