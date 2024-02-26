package main

import "fmt"
import "sort"

func main()  {
	fmt.Println("Hello from Slices Tuto")
	
	// var fruits[4] string

	// fruits[0]="Apple"
	// fruits[1]="Banana"
	// fruits[3]="Peach"

	// fmt.Println("Fruits list is",fruits)
	// fmt.Println("Length of array is",len(fruits))

	var veggies = []string{"potato","beans","mushrooms"}

	fmt.Println("Veggies List is",veggies)
	fmt.Printf("Type is %T\n",veggies)

	veggies = append(veggies,"Carrots","Cucumber")
	fmt.Println("Veggies List is",veggies)

	veggies = append(veggies[1:])
	fmt.Println("Veggies List is",veggies)
	veggies = append(veggies[1:3])
	fmt.Println("Veggies List is",veggies)

	highScores:= make([] int, 4)
	highScores[0]=145
	highScores[1]=245
	highScores[2]=345
	highScores[3]=445
	fmt.Println("Slice  is",highScores)  	// if you want to add one more value u need to use append. If you directly try to assign highScores[4]=545 then it wont allow cz the size is 4.(overflow it will lead to)
     
	sort.Ints(highScores)
	fmt.Println("Slice  is",highScores)
	fmt.Println(sort.IntsAreSorted(highScores))


	//Remove a value from slice based on index

	var courses = [] string{"reactjs","js","swift","c++","java"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)


}