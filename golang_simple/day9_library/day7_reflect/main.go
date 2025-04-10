package main

import (
	"fmt"
	"reflect"
)

/************ 空接口与反射 ************/
//func reflect_type(a interface{}) {
//	t := reflect.TypeOf(a)
//	fmt.Println("类型是：", t)
//	// kind()可以获取具体类型
//	k := t.Kind()
//	fmt.Println(k)
//	switch k {
//	case reflect.Float64:
//		fmt.Printf("a is float64\n")
//	case reflect.String:
//		fmt.Println("string")
//	}
//}
//
//func main() {
//	var x float64 = 3.4
//	reflect_type(x)
//}

//反射修改值
//func reflect_set_value(a interface{}) {
//	v := reflect.ValueOf(a)
//	k := v.Kind()
//	switch k {
//	case reflect.Float64:
//		// 反射修改值
//		v.SetFloat(6.9)
//		fmt.Println("a is ", v.Float())
//	case reflect.Ptr:
//		// Elem()获取地址指向的值
//		v.Elem().SetFloat(7.9)
//		fmt.Println("case:", v.Elem().Float())
//		// 地址
//		fmt.Println(v.Pointer())
//	}
//}
//
//func main() {
//	var x float64 = 3.4
//	// 反射认为下面是指针类型，不是float类型
//	reflect_set_value(&x)
//	fmt.Println("main:", x)
//}

/************ 结构体与反射 ************/
// 定义结构体
type User struct {
	Id   int
	Name string
	Age  int
}

// 绑方法
func (u User) Hello() {
	fmt.Println("Hello")
}

// 传入interface{}
func Poni(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("类型：", t)
	fmt.Println("字符串类型：", t.Name())
	// 获取值
	v := reflect.ValueOf(o)
	fmt.Println(v)
	// 可以获取所有属性
	// 获取结构体字段个数：t.NumField()
	for i := 0; i < t.NumField(); i++ {
		// 取每个字段
		f := t.Field(i)
		fmt.Printf("%s : %v", f.Name, f.Type)
		// 获取字段的值信息
		// Interface()：获取字段对应的值
		val := v.Field(i).Interface()
		fmt.Println("val :", val)
	}
	fmt.Println("================= 方法 ====================")
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name)
		fmt.Println(m.Type)
	}

}

func main() {
	u := User{1, "zs", 20}
	Poni(u)
}
