package main

import (
	"errors"
	"fmt"
)

func test() {
	//使用defer + recover 来捕获和处理异常
	defer func() {
		//recover()内置函数，可以捕获到异常
		if err := recover(); err != nil { //说明捕获到错误
			fmt.Println("err = ", err)
			//这里可以将错误信息发送给管理员...
			fmt.Println("发送邮件给admin@sohu.com")
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res = ", res)
}

//函数去读取一个配置文件init.conf的信息
//如果文件名传入不正确，我们就返回一个自定义的错误
func readConf(name string) (err error) {
	if name == "config.ini" {
		//读取...
		return nil
	} else {
		//返回一个自定义错误
		return errors.New("读取文件错误")
	}
}

func test02() {
	err := readConf("config.init")
	if err != nil {
		//如果读取文件发送错误，就输出这个错误，并终止程序
		panic(err)
	}
	fmt.Println("test02()继续执行...")
}

//Go 错误处理机制：defer,recover,error,panic
func main() {
	//测试
	// test()
	// for {
	// 	fmt.Println("main()下面的代码")
	// 	time.Sleep(time.Second)
	// }
	test02()
}
