package main

import "fmt"

func main() {

	var num1 float32 = 7.783
	fmt.Printf("The floating point is %g", num1) // we can use %g %f both for float
	var num2 = 15
	fmt.Printf("The binary value of num2  is %b", num2)
	// generic format specifier
	fmt.Printf("%v %v", num1, num2)

}
