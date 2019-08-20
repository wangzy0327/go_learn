package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱慕课网!" //UTF-8
	//每个中文占3字节
	fmt.Println(s)
	fmt.Println(len(s))
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s { //ch is a rune
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count:",
		utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	//rune 是 四个字节
	//使用range 遍历 string 得到 pos,rune对
	//使用len获得字节长度
	//使用[]byte获得字节
	for i, ch := range []rune(s) {
		fmt.Printf("(%d,%c)", i, ch)
	}

}
