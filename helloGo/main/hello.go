package main

import (
	"unsafe"
	"fmt"
)

func main(){
	fmt.Printf("Hello Go!\n")
	fmt.Println("Hello auto Promp!")
	fmt.Println("hello")
	fmt.Print("A", "B", "C")
	fmt.Println()
	var a = 10
	fmt.Printf( "%d", a )

	var name = "zhangsan1"
	var name2 string = "zhangsan2"
	name3 := "zhangsan3"

	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(name3)
	fmt.Printf("name1=%v name2=%v name3=%v \n", name, name2, name3)

	var num2 = 12
	fmt.Println(unsafe.Sizeof(num2))

	var number = 17
	// 原样输出
	fmt.Printf("%v\n", number)
	// 十进制输出
	fmt.Printf("%d\n", number)
	// 以八进制输出
	fmt.Printf("%o\n", number)
	// 以二进制输出
	fmt.Printf("%b\n", number)
	// 以十六进制输出
	fmt.Printf("%x\n", number)

	// for循环打印字符串里面的字符
	// 通过len来循环的，相当于打印的是ASCII码
	s := "你好 golang"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c)\t", s[i], s[i])
	}

	fmt.Println()

	// 通过rune打印的是 utf-8字符
	for index, v := range s {
		fmt.Printf("%v(%c)\n",index, v)
	}
	

	// rune类型
	s2 := "你好golang"


	
	// var byteS2 []rune = []rune(s2)
	var byteS2 []rune = []rune(s2)
	byteS2[0] = '我'
	fmt.Println(string(byteS2))

	// 字符串类型转换
	var i int = 20
	var f float64 = 12.456
	var t bool = true
	var b byte = 'a'
	str1 := fmt.Sprintf("%d", i)
	fmt.Printf("类型：%v-%T \n", str1, str1)

	str2 := fmt.Sprintf("%f", f)
	fmt.Printf("类型：%v-%T \n", str2, str2)

	str3 := fmt.Sprintf("%t", t)
	fmt.Printf("类型：%v-%T \n", str3, str3)

	str4 := fmt.Sprintf("%c", b)
	fmt.Printf("类型：%v-%T \n", str4, str4)

	//tip：在golang中，break可以不写，也能够跳出case，而不会执行其它的。
	extname := ".txt"
	switch extname {
		case ".html": {
			fmt.Println(".html")
		}
		case ".txt",".doc": {
			fmt.Println("传递来的是文档")
			break
		}
		case ".js": {
			fmt.Println(".js")
		}
		default: {
			fmt.Println("其它后缀")
		}
	}

	a1 := [...]int{1:1, 3:5}
	fmt.Println(a1)

}