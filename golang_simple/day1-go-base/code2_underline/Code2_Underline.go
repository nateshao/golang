package main

import "os"

func main() {
	buf := make([]byte, 1024)
	f, _ := os.Open("D:\\新建文本文档.txt")
	defer f.Close()
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}

/**
解释1：

    下划线意思是忽略这个变量.

    比如os.Open，返回值为*os.File，error

    普通写法是f,err := os.Open("xxxxxxx")

    如果此时不需要知道返回的错误值

    就可以用f, _ := os.Open("xxxxxx")

    如此则忽略了error变量
解释2：

    占位符，意思是那个位置本应赋给某个值，但是咱们不需要这个值。
    所以就把该值赋给下划线，意思是丢掉不要。
    这样编译器可以更好的优化，任何类型的单个值都可以丢给下划线。
    这种情况是占位用的，方法返回两个结果，而你只想要一个结果。
    那另一个就用 "_" 占位，而如果用变量的话，不使用，编译器是会报错的。
补充：

    import "database/sql"
    import _ "github.com/go-sql-driver/mysql"
第二个import就是不直接使用mysql包，只是执行一下这个包的init函数，把mysql的驱动注册到sql包里，然后程序里就可以使用sql包来访问mysql数据库了。




*/
