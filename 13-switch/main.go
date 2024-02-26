package main

import "fmt"
import "math/rand"
import "time"

func main() {
	fmt.Println("Switch Tuto")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
    fmt.Println("Value of dice is",diceNumber)
	switch diceNumber{
	    case 1:
			fmt.Println("Dice 1")
		case 2:
			fmt.Println("Dice 2")
		case 3:
			fmt.Println("Dice 3")
		case 4:
			fmt.Println("Dice 4")
		case 5:
			fmt.Println("Dice 5")
		case 6:
			fmt.Println("Dice 6")
		default:
			fmt.Println("Defaults")
	}

}
//fallthrough : next cases also will be printed(following cases)