package main

import "fmt"

/**
1. 匿名字段
Golang匿名字段 ：可以像字段成员那样访问匿名字段方法，编译器负责查找。
*/

//type User struct {
//	id   int
//	name string
//}
//
//type Manager struct {
//	User
//}
//
//func (self *User) ToString() string { // receiver = &(Manager.User)
//	return fmt.Sprintf("User: %p, %v", self, self)
//}
//
//func main() {
//	m := Manager{User{1, "Tom"}}
//	fmt.Printf("Manager: %p\n", &m)
//	fmt.Println(m.ToString())
//}

//

type User struct {
	id   int
	name string
}

type Manager struct {
	User
	title string
}

func (self *User) ToString() string {
	return fmt.Sprintf("User: %p, %v", self, self)
}

func (self *Manager) ToString() string {
	return fmt.Sprintf("Manager: %p, %v", self, self)
}

func main() {
	m := Manager{User{1, "Tom"}, "Administrator"}

	fmt.Println(m.ToString())

	fmt.Println(m.User.ToString(), 321)
}
