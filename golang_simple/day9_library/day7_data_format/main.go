package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string
	Hobby string
}

//func main() {
//	p := Person{"https://nateshao.gitlab.io/", "男"}
//	// 编码json
//	b, err := json.Marshal(p)
//	if err != nil {
//		fmt.Println("json err ", err)
//	}
//	fmt.Println(string(b))
//
//	// 格式化输出
//	b, err = json.MarshalIndent(p, "", "     ")
//	if err != nil {
//		fmt.Println("json err ", err)
//	}
//	fmt.Println(string(b))
//}

/********* 示例通过map生成json  *********/
type Person1 struct {
	Age       int    `json:"age,string"`
	Name      string `json:"name"`
	Niubility bool   `json:"niubility"`
}

//func main() {
//	// 假数据
//	b := []byte(`{"age":"18","name":"5lmh.com","marry":false}`)
//	var p Person1
//	err := json.Unmarshal(b, &p)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(p)
//}

func main() {
	// 假数据
	// int和float64都当float64
	b := []byte(`{"age":1.3,"name":"5lmh.com","marry":false}`)

	// 声明接口
	var i interface{}
	err := json.Unmarshal(b, &i)
	if err != nil {
		fmt.Println(err)
	}
	// 自动转到map
	fmt.Println(i)
	// 可以判断类型
	m := i.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case float64:
			fmt.Println(k, "是float64类型", vv)
		case string:
			fmt.Println(k, "是string类型", vv)
		default:
			fmt.Println("其他")
		}
	}
}
