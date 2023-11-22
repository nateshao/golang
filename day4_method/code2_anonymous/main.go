package main

import "fmt"

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

	fmt.Println(m.User.ToString())
}