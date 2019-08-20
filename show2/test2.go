package show2

import (
	"fmt"
	"test/show1"
)

func init(){
	fmt.Println("show2...")
}

func B(){
	show1.A()
	fmt.Println("B...")
}
