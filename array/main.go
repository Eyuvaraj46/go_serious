 package main
import "fmt"

 func main() {
	fmt.Println("welocome to the array program")
	
	
	var fruitlist [4]string //this is a format to declare an array
	
	fruitlist[0] = "apple"	
	fruitlist[1] = "banana" //get a space for [2]
	fruitlist[3] = "orange"

	fmt.Println("fruitlist is ", fruitlist)
	fmt .Println("fruitlist length is ", len(fruitlist))

	var veglist = [4]string{"potato", "tomato", "onion", "carrot"} //this is a format to declare an array with values
	fmt.Println("veglist is ", veglist)
	fmt.Println("veglist lemth is ",len(veglist))
 }

	