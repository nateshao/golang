package main

import (
	"fmt"
	"os"
	"time"
)

/**
1.自定义error
1.1. 抛异常和处理异常
1.1.1. 系统抛
*/

//func test1() {
//	a := [5]int{1, 2, 3, 4, 5}
//	a[1] = 123
//	fmt.Println(a)
//
//	index := 10
//	a[index] = 10
//	fmt.Println(a)
//}
//func getCircleArea(radius float32) (area float32) {
//	if radius < 0 {
//		// 自己抛
//		panic("半径不能为负")
//	}
//	return 3.14 * radius * radius
//}
//
//func test02() {
//	getCircleArea(-5)
//}
//
////
//func test03() {
//	// 延时执行匿名函数
//	// 延时到何时？（1）程序正常结束   （2）发生异常时
//	defer func() {
//		// recover() 复活 恢复
//		// 会返回程序为什么挂了
//		if err := recover(); err != nil {
//			fmt.Println(err)
//		}
//	}()
//	getCircleArea(-5)
//	fmt.Println("这里有没有执行")
//}
//func test04() {
//	test03()
//	fmt.Println("test04")
//}
//
//func main() {
//	test04()
//}

// 1.1.2. 返回异常

//func getCircleArea(radius float32) (area float32, err error) {
//	if radius < 0 {
//		// 构建个异常对象
//		err = errors.New("半径不能为负")
//		return
//	}
//	area = 3.14 * radius * radius
//	return
//}
//
//func main() {
//	area, err := getCircleArea(-5)
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(area)
//	}
//}

// 1.1.3. 自定义error：
type PathError struct {
	path       string
	op         string
	createTime string
	message    string
}

func (p *PathError) Error() string {
	return fmt.Sprintf("path=%s \nop=%s \ncreateTime=%s \nmessage=%s", p.path,
		p.op, p.createTime, p.message)
}

func Open(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return &PathError{
			path:       filename,
			op:         "read",
			message:    err.Error(),
			createTime: fmt.Sprintf("%v", time.Now()),
		}
	}
	defer file.Close()
	return nil
}
func main() {

	err := Open("/Users/5lmh/Desktop/go/src/test.txt")
	switch v := err.(type) {
	case *PathError:
		fmt.Println("get path error,", v)
	default:

	}
}
