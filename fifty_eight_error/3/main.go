package main

import (
	"log"
	"time"
)

//3.未使用的 import
//如果你 import一个包，但包中的变量、函数、接口和结构体一个都没有用到的话，将编译失败。可以使用 _下划线符号作为别名来忽略导入的包，从而避免编译错误，这只会执行 package 的 init()

// 错误示例
//import (
//"fmt"    // imported and not used: "fmt"
//"log"    // imported and not used: "log"
//"time"    // imported and not used: "time"
//)
//
//func main() {
//}

// 正确示例
// 可以使用 goimports 工具来注释或移除未使用到的包
import (
	_ "fmt"
)

func main() {
	_ = log.Println
	_ = time.Now
}
