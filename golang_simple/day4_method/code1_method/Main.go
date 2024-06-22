package main

import "fmt"

//结构体
type User struct {
	Name  string
	Email string
}

//方法
func (u *User) Notify() {
	fmt.Printf("%v : %v \n", u.Name, u.Email)
}

//func main() {
//	// 值类型调用方法
//	u1 := User{"golang", "golang@golang.com"}
//	u1.Notify()
//	// 指针类型调用方法
//	u2 := User{"go", "go@go.com"}
//	u3 := &u2
//	u3.Notify()
//}
/**
注意：当接受者是指针时，即使用值类型调用那么函数内部也是对指针的操作。

方法不过是一种特殊的函数，只需将其还原，就知道 receiver T 和 *T 的差别。
*/
//type Data struct {
//	x int
//}
//
//func (self Data) ValueTest() { // func ValueTest(self Data);
//	fmt.Printf("Value: %p\n", &self)
//}
//
//func (self *Data) PointerTest() { // func PointerTest(self *Data);
//	fmt.Printf("Pointer: %p\n", self)
//}
//
//func main() {
//	d := Data{}
//	p := &d
//	fmt.Printf("Data: %p\n", p)
//
//	d.ValueTest()   // ValueTest(d)
//	d.PointerTest() // PointerTest(&d)
//
//	p.ValueTest()   // ValueTest(*p)
//	p.PointerTest() // PointerTest(p)
//}

//普通函数与方法的区别（在接收者分别为值类型和指针类型的时候）

//1.普通函数
//接收值类型参数的函数
func valueIntTest(a int) int {
	return a + 10
}

//接收指针类型参数的函数
func pointerIntTest(a *int) int {
	return *a + 10
}

func structTestValue() {
	a := 2
	fmt.Println("valueIntTest:", valueIntTest(a))
	//函数的参数为值类型，则不能直接将指针作为参数传递
	//fmt.Println("valueIntTest:", valueIntTest(&a))
	//compile error: cannot use &a (type *int) as type int in function argument

	b := 5
	fmt.Println("pointerIntTest:", pointerIntTest(&b))
	//同样，当函数的参数为指针类型时，也不能直接将值类型作为参数传递
	//fmt.Println("pointerIntTest:", pointerIntTest(&b))
	//compile error:cannot use b (type int) as type *int in function argument
}

//2.方法
type PersonD struct {
	id   int
	name string
}

//接收者为值类型
func (p PersonD) valueShowName() {
	fmt.Println(p.name)
}

//接收者为指针类型
func (p *PersonD) pointShowName() {
	fmt.Println(p.name)
}

func structTestFunc() {
	//值类型调用方法
	personValue := PersonD{101, "hello world"}
	personValue.valueShowName()
	personValue.pointShowName()

	//指针类型调用方法
	personPointer := &PersonD{102, "hello golang"}
	personPointer.valueShowName()
	personPointer.pointShowName()

	//与普通函数不同，接收者为指针类型和值类型的方法，指针类型和值类型的变量均可相互调用
}

func main() {
	structTestValue()
	structTestFunc()
}
