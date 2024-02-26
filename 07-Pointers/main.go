package main

import "fmt"
// import "os"
// import "bufio"

func main()  {
	fmt.Println("Welcome to a class of Pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is", ptr);
	myNumber:= 23
	var ptr = &myNumber
	fmt.Println("Memory address is",ptr)
	fmt.Println("Value of actual pointer is",*ptr)
	*ptr = *ptr +2
	fmt.Println("Value of actual pointer is",*ptr)
}
