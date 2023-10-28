package main

/**
指针地址、指针类型和指针取值。
*/
//func main() {
//	//指针取值
//	a := 10
//	fmt.Println(a,&a)
//	b := &a // 取变量a的地址，将指针保存到b中
//	fmt.Printf("type of b:%T\n", b)
//	c := *b // 指针取值（根据指针去内存取值）
//	fmt.Printf("type of c:%T\n", c)
//	fmt.Printf("value of c:%v\n", c)
//}

/**
重要-----------------------------指针传值
*/
//func modify1(x int) {
//	x = 100
//}
//
//func modify2(x *int) {
//	*x = 100
//}
//
//func main() {
//	a := 10
//	modify1(a)
//	fmt.Println(a) // 10
//	modify2(&a)
//	fmt.Println(a) // 100
//}
//

//func main() {
//	var p *string
//	fmt.Println(p)
//	fmt.Printf("p的值是%v\n", p)
//	if p != nil {
//		fmt.Println("非空")
//	} else {
//		fmt.Println("空值")
//	}
//}

//  new和make
//func main() {
//	var a *int
//	*a = 100
//	fmt.Println(*a)
//
//	var b map[string]int
//	b["测试"] = 100
//	fmt.Println(b)
//	// 执行上面的代码会引发panic，为什么呢？ 在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。
//	//而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。要分配内存，就引出来今天的new和make。 Go语言中new和make是内建的两个函数，主要用来分配内存
//}

// new
//func main() {
//	a := new(int)
//	b := new(bool)
//	fmt.Printf("%T\n", a) // *int
//	fmt.Printf("%T\n", b) // *bool
//	fmt.Println(*a)       // 0
//	fmt.Println(*b)       // false
//}

//func main() {
//	var b map[string]int
//	b = make(map[string]int, 10)
//	b["测试"] = 100
//	fmt.Println(b)
//}

//指针小练习

//func main() {
//	var a int
//	fmt.Println(&a)
//	var p *int
//	p = &a
//	*p = 20
//	fmt.Println(a)
//}

// map基本使用

func main() {
}
