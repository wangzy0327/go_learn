package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func printArray(arr []int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	s1 := arr[2:]
	fmt.Println("arr[2:] = ", s1)
	s2 := arr[:]
	fmt.Println("arr[:] = ", s2)

	fmt.Println("After updateSlice (s1)")

	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateSlice (s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("Extending slice")
	arr[0], arr[2] = 0, 2
	fmt.Println("arr=", arr)
	s1 = arr[2:6]
	s2 = s1[3:5] // [s1[3],s1[4]]
	//fmt.Println(s1[4])
	//runtime error: index out of range

	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)

	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n",
		s1, len(s1), cap(s1))
	fmt.Printf("s2=%v,len(s2)=%d,cap(s2)=%d\n",
		s2, len(s2), cap(s2))

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)

	fmt.Println("s3,s4,s5 = ", s3, s4, s5)

	// s4 and s5 view different arr  no longer view arr
	//添加元素时如果超越cap，系统会重新分配更大的底层数组
	//由于值传递的关系，必须接受append的返回值（len和cap有可能会改变）
	fmt.Println("arr = ", arr)

	//fmt.Println(s1[3:7])
	//runtime error: slice bounds out of range

	//slice 可以向后扩展，不可以向前扩展 不可直接访问扩展的元素，但是可以看到
	//s[i]不可以超越len(s),向后扩展不可以超越底层数组cap(s)
	//	     slice
	//      ( ptr   len   cap)
	//[][][]||[][][][][][]|[][][][][][][]|

}
