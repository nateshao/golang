package main

import "fmt"

// 错误示例
type info struct {
	result int
}

func work() (int, error) {
	return 3, nil
}

// 正确示例
func main() {
	var data info
	var err error // err 需要预声明

	data.result, err = work()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("info: %+v\n", data)
}
