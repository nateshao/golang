package main

import "fmt"

// map基本使用

//func main() {
//	scoreMap := make(map[string]int, 8)
//	scoreMap["千羽"] = 90
//	scoreMap["燕燕"] = 90
//	fmt.Println(scoreMap)
//	fmt.Println(scoreMap["小明"])
//	fmt.Printf("type of a:%T\n",scoreMap)
//
//}

/**
  map也支持在声明的时候填充元素，例如：
*/

//func main() {
//	userInfo := map[string]string{
//		"username": "nateshao.gitlab.io",
//		"password": "123456",
//	}
//	fmt.Println(userInfo)
//}

//判断某个键是否存在
//Go语言中有个判断map中键是否存在的特殊写法，格式如下:
//
//value, ok := map[key]
func main() {
	scoreMap := make(map[string]int)
	scoreMap["千羽"] = 90
	scoreMap["千寻"] = 100
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["王五"] = 60
	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}
	/**
	遍历
	*/
	for s := range scoreMap {
		fmt.Println(s)
	}

	// 使用delete()函数删除键值对
	//使用delete()内建函数从map中删除一组键值对，delete()函数的格式如下：
	//
	//    delete(map, key)
	fmt.Println("-------------------- 使用delete()函数删除键值对 --------------------")
	delete(scoreMap, "小明")
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	fmt.Println("-------------------- 元素为map类型的切片 --------------------")

	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "王五"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "红旗大街"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}

	fmt.Println("-------------------- 值为切片类型的map --------------------")

	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)

	fmt.Println("-------------------- Go中Map的使用 --------------------")
	//直接创建初始化一个mao
	//var mapInitTemp = map[string]string {"xiaoli":"湖南", "xiaoliu":"天津"}
	//声明一个map类型变量,
	//map的key的类型是string，value的类型是string
	var mapTemp map[string]string
	//使用make函数初始化这个变量,并指定大小(也可以不指定)
	mapTemp = make(map[string]string, 10)
	//存储key ，value
	mapTemp["xiaoming"] = "北京"
	mapTemp["xiaowang"] = "河北"
	//根据key获取value,
	//如果key存在，则ok是true，否则是flase
	//v1用来接收key对应的value,当ok是false时，v1是nil
	v1, ok := mapTemp["xiaoming"]
	fmt.Println(ok, v1)
	//当key=xiaowang存在时打印value
	if v2, ok := mapTemp["xiaowang"]; ok {
		fmt.Println(v2)
	}
	//遍历map,打印key和value
	for k, v := range mapTemp {
		fmt.Println(k, v)
	}
	//删除map中的key
	delete(mapTemp, "xiaoming")
	//获取map的大小
	l := len(mapTemp)
	fmt.Println(l)

}
