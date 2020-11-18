package main

import (
	"fmt"
)


//Cat  定义一个Cat结构体
type Cat struct{
	Name string
	Age int
	Color string
}

func main(){
	var cat1 Cat
	fmt.Printf("cat1的地址:%p\n",&cat1)
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	fmt.Println("cat1=",cat1)

	fmt.Println("猫猫的信息如下：")
	fmt.Println("name = ",cat1.Name)
	fmt.Println("age = ",cat1.Age)
	fmt.Println("color = ",cat1.Color)
}