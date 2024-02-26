package main

import "fmt"
import "os"
import "io"
import "io/ioutil"
//Creating is a os operation ; reading is from ioutil
func main() {
  fmt.Println("Welcome to files in GoLang")
   content := "This needs to go into a file - baliga.s@northeastern.edu "
    file,err := os.Create("./myPersonalInfo.txt")

	if err != nil {
		panic(err)

	}
	
	length,err := io.WriteString(file,content)
	
	if err != nil {
		panic(err)

	}
	fmt.Println("Length is:",length)
	defer file.Close()
	readFile("./myPersonalInfo.txt")



}
func readFile(filename string) {
  databyte,err := ioutil.ReadFile(filename)
  if err != nil {
	panic(err)

   }
   fmt.Println("Text data inside the file is:",string(databyte))



}