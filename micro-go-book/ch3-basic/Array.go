package main

import "fmt"

func main() {

	//var classMates1 [3]string
	//
	//classMates1[0] = "小明"
	//classMates1[1] = "小红"
	//classMates1[2] = "小李"
	//fmt.Println(classMates1)
	//fmt.Println("The No.1 student is " + classMates1[0])
	//
	//classMates2 := [...]string{"小明", "小红", "小李"}
	//fmt.Println(classMates2)
	//
	//classMates3 := new([3]string)
	//classMates3[0] = "小明"
	//classMates3[1] = "小红"
	//classMates3[2] = "小李"
	//fmt.Println(*classMates3)

	// 写法1
	var arr1 [3]string
	arr1[0] = "nateshao1"
	arr1[1] = "nateshao2"
	arr1[2] = "nateshao3"
	fmt.Println("arr1 = ", arr1)

	// 写法2
	arr2 := [...]string{"小明", "小红", "小李"}
	fmt.Println("arr2 = ", arr2)

	// 写法3
	arr3 := new([3]string)
	arr3[0] = "xiaoming"
	arr3[1] = "xiaozi"
	arr3[2] = "xiaoshao"
	fmt.Println("arr3 = ", *arr3)
}
