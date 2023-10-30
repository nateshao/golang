package main

import "fmt"

type person struct {
	name string
	city string
	age  int8
}

func main() {
	var p1 person

	p1.name = "nateshao.gitlab.cn"
	p1.age = 20
	p1.city = "深圳"
	fmt.Println(p1)            // {nateshao.gitlab.cn 深圳 20}
	fmt.Printf("p1=%v\n", p1)  // p1={公众号：千羽的编程时光 北京 18}
	fmt.Printf("p1=%#v\n", p1) // p1=main.person{name:"公众号：千羽的编程时光", city:"北京", age:18}
	fmt.Println("------------------------- 匿名结构体 ------------------------------")

	var user struct {
		Name string
		Age  int
	}
	user.Name = "公众号：千羽的编程时光"
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
		name: "公众号：千羽的编程时光",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"公众号：千羽的编程时光", city:"北京", age:18}

	fmt.Println("------------------------- 使用值的列表初始化 ------------------------------")
	p8 := &person{
		"公众号：千羽的编程时光",
		"北京",
		18,
	}
	fmt.Printf("p8=%#v\n", p8) //p8=&main.person{name:"公众号：千羽的编程时光", city:"北京", age:18}

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
	p9 := newPerson("公众号：千羽的编程时光", "测试", 90)
	fmt.Printf("%#v\n", p9)
}
func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}
