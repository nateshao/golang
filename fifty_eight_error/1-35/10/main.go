package main

//10.map 容量
//在创建 map 类型的变量时可以指定容量，但不能像 slice 一样使用 cap() 来检测分配空间的大小：
//
//错误示例
//func main() {
//	m := make(map[string]int, 99)
//	println(cap(m))     // error: invalid argument m1 (type map[string]int) for cap
//}
