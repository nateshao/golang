package main

import (
	"fmt"
	"strings"
)

func main() {
	myString := "www.nateshao.com"

	// Option 1: (Recommended)
	if strings.HasPrefix(myString, "www") {
		fmt.Println("Hello to you too")
	} else {
		fmt.Println("Goodbye")
	}

}
