package main

// type Cat struct{}
//
// func (c Cat) Say() string { return "喵喵喵" }
//
// type Dog struct{}
//
// func (d Dog) Say() string { return "汪汪汪" }
//
//	func main() {
//		c := Cat{}
//		fmt.Println("猫:", c.Say())
//		d := Dog{}
//		fmt.Println("狗:", d.Say())
//	}
//type dog struct {
//}
//type cat struct {
//}
//
//func (d dog) say() {
//	fmt.Println("汪汪汪")
//}
//func (c cat) say() {
//	fmt.Println("喵喵喵")
//}
//func main() {
//	var x Sayer // 声明一个Sayer类型的变量x
//	a := cat{}  // 实例化一个cat
//	b := dog{}  // 实例化一个dog
//	x = a       // 可以把cat实例直接赋值给x
//	x.say()     // 喵喵喵
//	x = b       // 可以把dog实例直接赋值给x
//	x.say()     // 汪汪汪
//}

//type People interface {
//	Speak(string) string
//}
//
//type Student struct{}
//
//func (stu *Stduent) Speak(think string) (talk string) {
//	if think == "sb" {
//		talk = "你是个大帅比"
//	} else {
//		talk = "您好"
//	}
//	return
//}
//
//func main() {
//	var peo People = Student{}
//	think := "bitch"
//	fmt.Println(peo.Speak(think))
//}
