package main

import "fmt"

const LoginToken string = "abcdefgh"
func main()  {

	var userName string = "Shreya"
	fmt.Println(userName)
	fmt.Printf("Variable is of type: %T \n ",userName);
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n ",isLoggedIn);
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n ",smallVal);
	var smallFloat float32 = 255.5444
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n ",smallFloat);
	

	//default values and aliases
	var anotherVariable int
	fmt.Println(anotherVariable);
	fmt.Printf("Variable is of type: %T \n ",anotherVariable);


	//implicit ways of declaring varaibles
	var website = "shreyabaliga.me"
	fmt.Println(website); //works fine
	//But when u try to change again it wont allow  


	//no var style   -- These kind are Allowed only inside method 
	numberOfUser:= 30000.0
	fmt.Println(numberOfUser);


	fmt.Println(LoginToken);
	fmt.Printf("Is of type: %T\n",LoginToken);

 	
	
}

