package main

import "fmt"
// import "sort"
// There is no inheritance, concept of super in GoLang
func main()  {
	fmt.Println("Hello from Structs Tuto")
	
	Shreya := User{"Shreya_Baliga","baliga.s@northeastern.edu",true,24}
	fmt.Println(Shreya)
	fmt.Printf("Shreya's details are %+v\n",Shreya)
	fmt.Printf("Shreya's details are :Name is %v and email is %v and age is %v\n",Shreya.Name,Shreya.Email,Shreya.age)
}

type User struct{
	Name string
	Email string
	Status bool
	age int

}