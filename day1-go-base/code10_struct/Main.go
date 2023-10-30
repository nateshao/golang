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
	fmt.Printf("p1=%v\n", p1)  // p1={pprof.cn 北京 18}
	fmt.Printf("p1=%#v\n", p1) // p1=main.person{name:"pprof.cn", city:"北京", age:18}

}
