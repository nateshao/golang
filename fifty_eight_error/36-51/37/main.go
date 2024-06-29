package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// 主动关闭连接
func main() {
	tr := http.Transport{DisableKeepAlives: true}
	client := http.Client{Transport: &tr}

	resp, err := client.Get("https://golang.google.cn/")
	if resp != nil {
		defer resp.Body.Close()
	}
	checkError(err)

	fmt.Println(resp.StatusCode) // 200

	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	fmt.Println(len(string(body)))
}
