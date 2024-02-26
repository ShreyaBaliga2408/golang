package main

import "fmt"

func main() {
	fmt.Println("Functions Tuto")
	greetings()
	fmt.Println("Function executed")
	greeter()
    fmt.Println("Function executed")

	result:= addition(3,5)
	fmt.Println(result)

	proResult:= prodAdd(2,3,4,5,6,6,7)
	fmt.Println(proResult)
}
func addition(value1 int,value2 int) int {
	return value1+value2
}

func prodAdd(values ...int) int {
	total:=0
	for _,val := range values {
		total+=val
	}
	return total
}
func greeter() {
	fmt.Println("Another Method written")
}
func greetings() {
	fmt.Println("Hi from Shreya")
}