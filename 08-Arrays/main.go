package main

import "fmt"

func main()  {
	fmt.Println("Hello from Arrays Tuto")
	
	var fruits[4] string

	fruits[0]="Apple"
	fruits[1]="Banana"
	fruits[3]="Peach"

	fmt.Println("Fruits list is",fruits)
	fmt.Println("Length of array is",len(fruits))

	var veggies = [3]string{"potato","beans","mushrooms"}
	fmt.Println("Veggies List is",veggies)
}