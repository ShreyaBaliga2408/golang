package main

import "fmt"
import "os"
import "bufio"

func main()  {
	welcome:= "Welcome to user Input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for our Pizza:")

	// comma ok || error err
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating",input);
	fmt.Printf("Type of rating is: %T\n",input);
	
}

