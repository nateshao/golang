package main

import (
	"encoding/json"
	"fmt"
)

type MyData struct {
	a int
	b string
}

func main() {
	in := MyData{1, "two"}
	fmt.Printf("%#v\n", in) // main.MyData{One:1, two:"two"}

	encoded, _ := json.Marshal(in)
	fmt.Println(string(encoded)) // {"One":1}    // 私有字段 two 被忽略了

	var out MyData
	json.Unmarshal(encoded, &out)
	fmt.Printf("%#v\n", out) // main.MyData{One:1, two:""}
}
