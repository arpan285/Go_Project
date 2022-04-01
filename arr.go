package main

import "fmt"

func main() {

	var a int
	//var arr [100]int

	fmt.Printf("Enter a number: ")
	fmt.Scanf("%d", &a)
	//fmt.Println("Your numbers is:", a)

	var arr = make([]int, a)

	for i := 0; i < a; i++ {
		fmt.Printf("Enter %dth element: ", i)
		fmt.Scanf("%d", &arr[i])
	}

	fmt.Println("Your array is: ", arr)
	/*for i := 0; i < a; i++ {
		fmt.Printf("%d", arr[i])
	}*/
}
