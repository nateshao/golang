package main

import "fmt"

/**
fmt包实现了类似C语言printf和scanf的格式化I/O。主要分为向外输出内容和获取输入内容两大部分。


*/
func main() {
	//fmt.Print("在终端打印该信息。")
	//name := "枯藤"
	//fmt.Printf("我是：%s\n", name)
	//fmt.Println("在终端打印单独一行显示")
	fmt.Println("---------- Fprint -------------")

	s1 := fmt.Sprint("枯藤")
	name := "枯藤"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("枯藤")
	fmt.Println(s1, s2, s3)

	fmt.Println("---------- Errorf -------------")
	/**
	Errorf函数根据format参数生成格式化字符串并返回一个包含该字符串的错误。
	*/
	err := fmt.Errorf("这是一个错误")
	fmt.Println(err)

	fmt.Println("---------- 格式化占位符 -------------")
	fmt.Printf("%v\n", 100)
	fmt.Printf("%v\n", false)
	o := struct{ name string }{"枯藤"}
	fmt.Printf("%v\n", o)
	fmt.Printf("%#v\n", o)
	fmt.Printf("%T\n", o)
	fmt.Printf("100%%\n")
}
