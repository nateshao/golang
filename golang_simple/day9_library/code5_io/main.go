package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	writeFile()
	readFile()
	copyFile()
}

/**
打开和关闭文件
*/
func openAndCloseFile() {

	// 只读方式打开当前目录下的main.go文件
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	// 关闭文件
	file.Close()
}

/**
写文件
*/
func writeFile() {
	file, err := os.Create("./writeFile.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	for i := 0; i < 5; i++ {
		file.WriteString("ab\n")
		file.Write([]byte("cd\n"))
	}
}

/**
读文件
*/
func readFile() {
	file, err := os.Open("./writeFile.txt")
	if err != nil {
		fmt.Println("open file err :", err)
		return
	}
	defer file.Close()
	// 定义接收文件读取的字节数组
	var buf [128]byte
	var content []byte
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			// 读取结束
			break
		}
		if err != nil {
			fmt.Println("read file err ", err)
			return
		}
		content = append(content, buf[:n]...)
	}
	fmt.Println(string(content))
}

/**
拷贝文件
*/

func copyFile() {
	// 打开源文件
	srcFile, err := os.Open("./writeFile.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 创建新文件
	dstFile, err2 := os.Create("./abc2.txt")
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	// 缓冲读取
	buf := make([]byte, 1024)
	for {
		// 从源文件读数据
		n, err := srcFile.Read(buf)
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		//写出去
		dstFile.Write(buf[:n])
	}
	srcFile.Close()
	dstFile.Close()
}

/**
bufio
*/
func bufioDemo() {

}
