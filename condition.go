package main

import "fmt"

func main() {

	var a, b int
	fmt.Println("Enter two nubers: ")
	fmt.Scanf("%v %v", &a, &b)
	fmt.Println("Your numbers are:", a, "and", b)
	if a > b {
		fmt.Println("a is greater than b")
	} else {
		fmt.Println("b is greater than a")
	}
}
