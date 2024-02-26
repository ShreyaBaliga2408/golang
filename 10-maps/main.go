package main

import "fmt"
// import "sort"

func main()  {
	fmt.Println("Hello from Maps Tuto")
	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages:" , languages)
	fmt.Println("JS languages:" , languages["JS"]);

	delete(languages, "RB")
	fmt.Println("List of all languages:" , languages)
   
	//loops in maps;
	for key,value := range languages {
		fmt.Printf("For key %v, value is %v\n" , key,value)
	}
	
	
}