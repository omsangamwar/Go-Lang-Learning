package main

import "fmt"

func main() {
	var year int = 0

	fmt.Printf("Enter Number")
	fmt.Scanf("%d", &year)
	if year != 0 {
		if (year%4 == 0 && year%100 != 0) || (year%4 == 0 && year%100 == 0 && year%400 == 0) {
			fmt.Printf("Entered Year is  Leap Year")
		} else {
			fmt.Printf("Entered Year is Leap Year")
		}
	}

}
