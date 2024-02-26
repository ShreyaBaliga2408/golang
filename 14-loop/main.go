package main

import "fmt"

func main()  {
	fmt.Println("Hello from Loops Tuto")

   days := []string { "Sun" ,"Mon","Tue","Wed","thu","Fri","Sat"}
   fmt.Println(days)

//    for i:=0; i< len(days); i++ {
// 	fmt.Println(days[i])
//    }	
//    for i := range days {
// 	fmt.Println(days[i])
//    }

//    for index,day := range days {
// 	 fmt.Printf("Index is %v and value at that index is %v\n",index,day)
//    }
//    for _,day := range days {
// 	fmt.Printf("Index is  and value at that index is %v\n",day)
//   }

  rogueValue := 1
  for rogueValue < 10 {
	fmt.Println(rogueValue)
	rogueValue++;

	if rogueValue==5 {
        break;
	}
	if rogueValue==2 {
		rogueValue++;
		continue;
		goto lco
		
	}
	  
  }
  lco:
  fmt.Println("Jumping statements")
}
