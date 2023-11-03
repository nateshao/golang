package main

import (
	"fmt"
	"reflect"
	"testing"
)

//func main() {
//
//}
// 切割测试
func TestSplit(t *testing.T) { // 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
	got := Split("a:b:c", ":") // 程序输出的结果
	fmt.Println(got)
	want := []string{"a", "b", "c"} // 期望的结果
	fmt.Println(want)
	if !reflect.DeepEqual(want, got) { // 因为slice不能比较直接，借助反射包中的方法比较
		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
	}

}

func TestMoreSplit(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("excepted:%v, got:%v", want, got)
	}

	//got := Split("abcd", "bc")
	//want := []string{"a", "d"}
	//if !reflect.DeepEqual(want, got) {
	//	t.Errorf("excepted:%v, got:%v", want, got)
	//}

}

func TestChineseSplit(t *testing.T) {

}
