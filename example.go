package main

import "fmt"

func main() {

	fmt.Println("1+1=", 1+1)
	var a int = 21
	var b string = " Arpan"
	s := 1000 //inferred
	var x float64 = 3.4e+39
	var txt string = " High 5"

	fmt.Println("age=", a)
	fmt.Println("Name=", b)
	fmt.Println("Salary=", s)
	fmt.Println(x)
	fmt.Printf("Type= %T \n & Value=%v", txt, txt)
}
