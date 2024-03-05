package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "hello World"
	var subStr string = "Wor"

	if strings.Contains(str, subStr) == true {
		fmt.Printf("String (%s) contains sub-string (%s)", str, subStr)
	} else {
		fmt.Printf("String (%s)  does not contains sub-string (%s)", str, subStr)
	}

	x := "I am om"
	fmt.Println("Who are you ??", strings.ToUpper(x))
	fmt.Println("Who are you ??", strings.ToLower(x))

}
