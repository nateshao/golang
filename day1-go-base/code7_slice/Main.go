package main

import "fmt"

/**
切片Slice
需要说明，slice 并不是数组或数组指针。它通过内部指针和相关属性引用数组片段，以实现变长方案。

    1. 切片：切片是数组的一个引用，因此切片是引用类型。但自身是结构体，值拷贝传递。
    2. 切片的长度可以改变，因此，切片是一个可变的数组。
    3. 切片遍历方式和数组一样，可以用len()求长度。表示可用元素数量，读写操作不能超过该限制。
    4. cap可以求出slice最大扩张容量，不能超出数组限制。0 <= len(slice) <= len(array)，其中array是slice引用的数组。
    5. 切片的定义：var 变量名 []类型，比如 var str []string  var arr []int。
    6. 如果 slice == nil，那么 len、cap 结果都等于 0。
*/
//func main() {
//	// 声明切片
//	var s1 []int
//	if s1 == nil {
//		fmt.Println("是空")
//	} else {
//		fmt.Println("不是空")
//	}
//
//	// 2.:=
//	s2 := []int{}
//	// 3.make()
//	var s3 []int = make([]int, 0)
//	fmt.Println(s1, s2, s3)
//	// 4.初始化赋值
//	var s4 []int = make([]int, 0, 0)
//	fmt.Println(s4)
//	s5 := []int{1, 2, 3}
//	fmt.Println(s5)
//	// 5.从数组切片
//	arr := [5]int{1, 2, 3, 4, 5}
//	var s6 []int
//	// 前包后不包
//	s6 = arr[1:4]
//	fmt.Println(s6)
//}

/********************************  *************************************/

var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var slice0 []int = arr[2:8]
var slice1 []int = arr[0:6]        //可以简写为 var slice []int = arr[:end]
var slice2 []int = arr[5:10]       //可以简写为 var slice[]int = arr[start:]
var slice3 []int = arr[0:len(arr)] //var slice []int = arr[:]
var slice4 = arr[:len(arr)-1]      //去掉切片的最后一个元素
func main() {
	fmt.Printf("全局变量：arr %v\n", arr)
	fmt.Printf("全局变量：slice0 %v\n", slice0)
	fmt.Printf("全局变量：slice1 %v\n", slice1)
	fmt.Printf("全局变量：slice2 %v\n", slice2)
	fmt.Printf("全局变量：slice3 %v\n", slice3)
	fmt.Printf("全局变量：slice4 %v\n", slice4)
	fmt.Printf("-----------------------------------\n")
	arr2 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	slice5 := arr[2:8]
	slice6 := arr[0:6]         //可以简写为 slice := arr[:end]
	slice7 := arr[5:10]        //可以简写为 slice := arr[start:]
	slice8 := arr[0:len(arr)]  //slice := arr[:]
	slice9 := arr[:len(arr)-1] //去掉切片的最后一个元素
	fmt.Printf("局部变量： arr2 %v\n", arr2)
	fmt.Printf("局部变量： slice5 %v\n", slice5)
	fmt.Printf("局部变量： slice6 %v\n", slice6)
	fmt.Printf("局部变量： slice7 %v\n", slice7)
	fmt.Printf("局部变量： slice8 %v\n", slice8)
	fmt.Printf("局部变量： slice9 %v\n", slice9)
}

//通过make来创建切片
//var slice []type = make([]type, len)
//slice  := make([]type, len)
//slice  := make([]type, len, cap)
