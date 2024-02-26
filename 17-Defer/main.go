package main

import "fmt"
//When u have multiple defer - While Printing LIFO is followed 
func main() {

	fmt.Println("Hello from Defer Tuto")
	defer fmt.Println("Hello World")
	defer fmt.Println("Hi")
	defer fmt.Println("Hey")
	defer fmt.Println("Namaste")
	fmt.Println("Hey from Shreya")

	//OUTPUT:
	// Hello from Defer Tuto
	// Hey from Shreya
	// Namaste
	// Hey
	// Hi
	// Hello World

myDeferFunction()

//OUTPUT:
// Hello from Defer Tuto
// Hey from Shreya
// 4
// 3
// 2
// 1
// 0
// Namaste
// Hey
// Hi

}

func myDeferFunction() {
	for i:=0; i < 5; i++ {
		defer fmt.Println(i)
	}
}