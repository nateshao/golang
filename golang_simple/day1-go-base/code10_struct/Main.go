package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	name string
	city string
	age  int8
}

func main() {
	var p1 person

	p1.name = "nateshao.gitlab.io"
	p1.age = 20
	p1.city = "深圳"
	fmt.Println(p1)            // {nateshao.gitlab.cn 深圳 20}
	fmt.Printf("p1=%v\n", p1)  // p1={公众号：程序员千羽 北京 18}
	fmt.Printf("p1=%#v\n", p1) // p1=main.person{name:"公众号：程序员千羽", city:"北京", age:18}
	fmt.Println("------------------------- 匿名结构体 ------------------------------")

	var user struct {
		Name string
		Age  int
	}
	user.Name = "公众号：程序员千羽"
	user.Age = 18
	fmt.Printf("%#v\n", user)
	// 创建指针类型结构体
	var p2 = new(person)
	fmt.Printf("%T\n", p2)     //*main.person
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"", city:"", age:0}

	fmt.Println("------------------------- 创建指针类型结构体 ------------------------------")
	//var p2 = new(person)
	//p2.name = "测试"
	//p2.age = 18
	//p2.city = "北京"
	//fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"测试", city:"北京", age:18}

	//

	fmt.Println("------------------------- 取结构体的地址实例化 ------------------------------")
	p3 := &person{}
	fmt.Printf("%T\n", p3)     //*main.person
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
	p3.name = "博客"
	p3.age = 30
	p3.city = "成都"
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"博客", city:"成都", age:30}

	fmt.Println("------------------------- 结构体初始化 ------------------------------")

	var p4 person
	fmt.Printf("p4=%#v\n", p4) //p4=main.person{name:"", city:"", age:0}

	fmt.Println("------------------------- 使用键值对初始化 ------------------------------")
	p5 := person{
		name: "公众号：程序员千羽",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"公众号：程序员千羽", city:"北京", age:18}

	fmt.Println("------------------------- 使用值的列表初始化 ------------------------------")
	p8 := &person{
		"公众号：程序员千羽",
		"北京",
		18,
	}
	fmt.Printf("p8=%#v\n", p8) //p8=&main.person{name:"公众号：程序员千羽", city:"北京", age:18}

	fmt.Println("------------------------- 结构体内存布局 ------------------------------")

	type test struct {
		a int8
		b int8
		c int8
		d int8
	}
	n := test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)

	fmt.Println("------------------------- 构造函数 ------------------------------")
	p9 := newPerson("公众号：程序员千羽", "测试", 90)
	fmt.Printf("%#v\n", p9)

	fmt.Println("------------------------- 结构体与JSON序列化 ------------------------------")
	c := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := &Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		}
		c.Students = append(c.Students, stu)
	}
	//JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)
	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)
}

// Student 学生
type Student struct {
	ID     int
	Gender string
	Name   string
}

// Class 班级
type Class struct {
	Title    string
	Students []*Student
}

func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

// Person 结构体
type Person struct {
	name string
	age  int8
}

// Dream Person做梦的方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

// SetAge 设置p的年龄
// 使用指针接收者
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}
