package main

import "fmt"
// We dont have classes in GoLang but we have structs in GoLang; So functions inside structs are called Methods.
func main()  {
	fmt.Println("Hello from Methods Tuto")
	
	Shreya := User{"Shreya_Baliga","baliga.s@northeastern.edu",true,24}
	fmt.Println(Shreya)
	fmt.Printf("Shreya's details are %+v\n",Shreya)
	fmt.Printf("Shreya's details are :Name is %v and email is %v and age is %v\n",Shreya.Name,Shreya.Email,Shreya.age)
	Shreya.getStatus()
	Shreya.newMail()

}

type User struct{
	Name string
	Email string
	Status bool
	age int

}

func(u User) getStatus() {

	fmt.Println("Is user active? : ", u.Status)


}

func(u User) newMail() {
	u.Email = "test@go.dev"
	fmt.Println("New Email of user is :" , u.Email)
}