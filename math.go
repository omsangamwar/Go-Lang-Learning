package main

import (
	"fmt"
	"math"
)

func main() {
	var a int = 225
	var b float32
	b = float32((math.Sqrt(float64(a))))

	fmt.Printf("the square root of 123 is :%f", b)

}

//math.Sqrt takes only float 64 arguments
