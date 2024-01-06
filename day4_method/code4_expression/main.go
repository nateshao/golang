package main

import "fmt"

/**
1. 表达式
Golang 表达式 ：根据调用者不同，方法分为两种表现形式:

    instance.method(args...) ---> <type>.func(instance, args...)
前者称为 method value，后者 method expression。

两者都可像普通函数那样赋值和传参，区别在于 method value 绑定实例，而 method expression 则须显式传参。
*/

//type User struct {
//	id   int
//	name string
//}
//
//func (self *User) Test() {
//	fmt.Printf("%p, %v\n", self, self)
//}

/**
%p 格式化占位符输出变量 self 的内存地址（指针的十六进制表示）。
%v 格式化占位符输出变量 self 所指向的值，以默认的格式进行显示。
*/

//func main() {
//	user := User{1, "程序员千羽"}
//	user.Test()
//
//	mValue := user.Test
//	mValue() // 隐式传递 receiver
//
//	mExpression := (*User).Test
//	mExpression(&user) // 显式传递 receiver
//}

/******************** 2 **********************/

type User struct {
	id   int
	name string
}

func (self User) Test() {
	fmt.Println(self)
}

func main() {
	u := User{1, "Tom"}
	mValue := u.Test // 立即复制 receiver，因为不是指针类型，不受后续修改影响。

	u.id, u.name = 2, "Jack"
	u.Test()

	mValue()
}
