package main

import (
	"fmt"
	"reflect"
	"test/show1"
	"test/show2"
	"unsafe"
)

type 慕课网 int32

func main() {
	show2.B()
	show1.A()
	var i bool = false
	fmt.Println(unsafe.Sizeof(i))
	var j int32 = 22
	fmt.Println(unsafe.Sizeof(j))
	var k rune = 32
	fmt.Println(unsafe.Sizeof(k))
	var d 慕课网
	fmt.Println(reflect.TypeOf(d))
	fmt.Println("hello GO!")
}
