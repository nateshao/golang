package main

import "fmt"

/**
切片Slice
需要说明，code1_slice 并不是数组或数组指针。它通过内部指针和相关属性引用数组片段，以实现变长方案。

    1. 切片：切片是数组的一个引用，因此切片是引用类型。但自身是结构体，值拷贝传递。
    2. 切片的长度可以改变，因此，切片是一个可变的数组。
    3. 切片遍历方式和数组一样，可以用len()求长度。表示可用元素数量，读写操作不能超过该限制。
    4. cap可以求出slice最大扩张容量，不能超出数组限制。0 <= len(code1_slice) <= len(array)，其中array是slice引用的数组。
    5. 切片的定义：var 变量名 []类型，比如 var str []string  var arr []int。
    6. 如果 code1_slice == nil，那么 len、cap 结果都等于 0。
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

//var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
//var slice0 []int = arr[2:8]
//var slice1 []int = arr[0:6]        //可以简写为 var code1_slice []int = arr[:end]
//var slice2 []int = arr[5:10]       //可以简写为 var code1_slice[]int = arr[start:]
//var slice3 []int = arr[0:len(arr)] //var code1_slice []int = arr[:]
//var slice4 = arr[:len(arr)-1]      //去掉切片的最后一个元素
//func main() {
//	fmt.Printf("全局变量：arr %v\n", arr)
//	fmt.Printf("全局变量：slice0 %v\n", slice0)
//	fmt.Printf("全局变量：slice1 %v\n", slice1)
//	fmt.Printf("全局变量：slice2 %v\n", slice2)
//	fmt.Printf("全局变量：slice3 %v\n", slice3)
//	fmt.Printf("全局变量：slice4 %v\n", slice4)
//	fmt.Printf("-----------------------------------\n")
//	arr2 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
//	slice5 := arr[2:8]
//	slice6 := arr[0:6]         //可以简写为 code1_slice := arr[:end]
//	slice7 := arr[5:10]        //可以简写为 code1_slice := arr[start:]
//	slice8 := arr[0:len(arr)]  //code1_slice := arr[:]
//	slice9 := arr[:len(arr)-1] //去掉切片的最后一个元素
//	fmt.Printf("局部变量： arr2 %v\n", arr2)
//	fmt.Printf("局部变量： slice5 %v\n", slice5)
//	fmt.Printf("局部变量： slice6 %v\n", slice6)
//	fmt.Printf("局部变量： slice7 %v\n", slice7)
//	fmt.Printf("局部变量： slice8 %v\n", slice8)
//	fmt.Printf("局部变量： slice9 %v\n", slice9)
//}

//通过make来创建切片
//var code1_slice []type = make([]type, len)
//code1_slice  := make([]type, len)
//code1_slice  := make([]type, len, cap)

/********************************  *************************************/

//func main() {
//	data := [...]int{0, 1, 2, 3, 4, 5}
//
//	s := data[2:4]
//	fmt.Println(s)
//	s[0] += 100
//	s[1] += 200
//
//	fmt.Println(s)
//	fmt.Println(data)
//}

/********************************  *************************************/
//func main() {
//	s1 := []int{0, 1, 2, 3, 8: 100} // 通过初始化表达式构造，可使用索引号。
//	fmt.Println(s1, len(s1), cap(s1))
//
//	s2 := make([]int, 6, 8) // 使用 make 创建，指定 len 和 cap 值。
//	fmt.Println(s2, len(s2), cap(s2))
//
//	s3 := make([]int, 6) // 省略 cap，相当于 cap = len。
//	fmt.Println(s3, len(s3), cap(s3))
//}

/********************************  *************************************/
//func main() {
//	s := []int{0, 1, 2, 3}
//	fmt.Println(s)
//	p := &s[2] // *int, 获取底层数组元素指针。
//	//println(p)
//	*p += 100
//
//	fmt.Println(s)
//}

/********************************  用append内置函数操作切片（切片追加） *************************************/

//func main() {
//
//	var a = []int{1, 2, 3}
//	fmt.Printf("code1_slice a : %v\n", a)
//	var b = []int{4, 5, 6}
//	fmt.Printf("code1_slice b : %v\n", b)
//	c := append(a, b...)
//	fmt.Printf("code1_slice c : %v\n", c)
//	d := append(c, 7)
//	fmt.Printf("code1_slice d : %v\n", d)
//	e := append(d, 8, 9, 10)
//	fmt.Printf("code1_slice e : %v\n", e)
//
//}

/********************************  *************************************/
//func main() {
//
//	s1 := make([]int, 0, 5)
//	fmt.Printf("%p\n", &s1)
//
//	s2 := append(s1, 1)
//	fmt.Printf("%p\n", &s2)
//
//	fmt.Println(s1, s2)
//
//}

/******************************** 超出原 code1_slice.cap 限制，就会重新分配底层数组，即便原数组并未填满。 *************************************/

//import (
//	"fmt"
//)
//
//func main() {
//
//	data := [...]int{0, 1, 2, 3, 4, 10: 0}
//	s := data[:2:3]
//	s = append(s, 100, 200) // 一次 append 两个值，超出 s.cap 限制。
//
//	fmt.Println(s, data)         // 重新分配底层数组，与原数组无关。
//	fmt.Println(&s[0], &data[0]) // 比对底层数组起始指针。
//
//}

/******************************** slice中cap重新分配规律  *************************************/
//import (
//	"fmt"
//)
//
//func main() {
//
//	s := make([]int, 0, 1)
//	fmt.Println(s)
//	c := cap(s)
//	fmt.Println(c)
//
//	for i := 0; i < 50; i++ {
//		s = append(s, i)
//		if n := cap(s); n > c {
//			fmt.Printf("cap: %d -> %d\n", c, n)
//			c = n
//		}
//	}
//
//}

/********************************  切片拷贝 *************************************/
//func main() {
//
//	s1 := []int{1, 2, 3, 4, 5}
//	fmt.Printf("code1_slice s1 : %v\n", s1)
//	s2 := make([]int, 10)
//	fmt.Printf("code1_slice s2 : %v\n", s2)
//	copy(s2, s1)
//	fmt.Printf("copied code1_slice s1 : %v\n", s1)
//	fmt.Printf("copied code1_slice s2 : %v\n", s2)
//	s3 := []int{1, 2, 3}
//	fmt.Printf("code1_slice s3 : %v\n", s3)
//	s3 = append(s3, s2...)
//	fmt.Printf("appended code1_slice s3 : %v\n", s3)
//	s3 = append(s3, 4, 5, 99)
//	fmt.Printf("last code1_slice s3 : %v\n", s3)
//}

/******************************** copy ：函数 copy 在两个 code1_slice 间复制数据，复制长度以 len 小的为准。两个 code1_slice 可指向同一底层数组，允许元素区间重叠。 *************************************/

//func main() {
//
//	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
//	fmt.Println("array data : ", data)
//	s1 := data[8:]
//	s2 := data[:5]
//	fmt.Printf("code1_slice s1 : %v\n", s1)
//	fmt.Printf("code1_slice s2 : %v\n", s2)
//	copy(s2, s1)
//	fmt.Printf("copied code1_slice s1 : %v\n", s1)
//	fmt.Printf("copied code1_slice s2 : %v\n", s2)
//	fmt.Println("last array data : ", data)
//
//}

/******************************** slice遍历：  *************************************/
//func main() {
//
//	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
//	code1_slice := data[:]
//	fmt.Println(code1_slice)
//	for index, value := range code1_slice {
//		fmt.Printf("inde : %v , value : %v\n", index, value)
//	}
//
//}

/******************************** 含有中文字符串： *************************************/
//func main() {
//	str := "你好，世界！hello world！"
//	s := []rune(str)
//	s[3] = '够'
//	s[4] = '浪'
//	s[12] = 'g'
//	s = s[:14]
//	str = string(s)
//	fmt.Println(str)
//}

/********************************  *************************************/
//func main() {
//	code1_slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
//	d1 := code1_slice[6:8]
//	fmt.Println(d1, len(d1), cap(d1))
//	d2 := code1_slice[:6:8]
//	fmt.Println(d2, len(d2), cap(d2))
//}

/********************************  *************************************/

func main() {
	var numbers = make([]int, 3, 5)

	ints := append(numbers, 3)
	fmt.Println(ints)
}
