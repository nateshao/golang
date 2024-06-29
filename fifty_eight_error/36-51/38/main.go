package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

//// 将 decode 的值转为 int 使用
//func main() {
//	var data = []byte(`{"status": 200}`)
//	var result map[string]interface{}
//
//	if err := json.Unmarshal(data, &result); err != nil {
//		log.Fatalln(err)
//	}
//
//	var status = uint64(result["status"].(float64))
//	fmt.Println("Status value: ", status)
//}

//type Person struct {
//	Name string `json:"name"`
//	Age  int    `json:"age"`
//}
//
//func main() {
//	jsonStr := `{"name":"John", "age":30}`
//
//	var p Person
//	err := json.Unmarshal([]byte(jsonStr), &p)
//	if err != nil {
//		log.Fatalf("Error unmarshaling JSON: %v", err)
//	}
//
//	fmt.Printf("Person: %+v\n", p) // 输出: Person: {Name:John Age:30}
//}
//
//// struct 中指定字段类型
//func main() {
//	var data = []byte(`{"status": 200}`)
//	var result struct {
//		Status uint64 `json:"status"`
//	}
//
//	err := json.NewDecoder(bytes.NewReader(data)).Decode(&result)
//	checkError(err)
//	fmt.Printf("Result: %+v", result)
//}

//// 指定字段类型
//func main() {
//	var data = []byte(`{"status": 200}`)
//	var result map[string]interface{}
//
//	var decoder = json.NewDecoder(bytes.NewReader(data))
//	decoder.UseNumber()
//
//	if err := decoder.Decode(&result); err != nil {
//		log.Fatalln(err)
//	}
//
//	var status, _ = result["status"].(json.Number).Int64()
//	fmt.Println("Status value: ", status)
//}

// struct 中指定字段类型
func main() {
	var data = []byte(`{"status": 200}`)
	var result struct {
		Status uint64 `json:"status"`
	}

	err := json.NewDecoder(bytes.NewReader(data)).Decode(&result)
	checkError(err)
	fmt.Printf("Result: %+v", result)
}
