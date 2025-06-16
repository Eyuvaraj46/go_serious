package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitlist = []string{"apple", "banana", "papaya", "pinapple"} //infinite index to add infinite element
	fruitlist = append(fruitlist, "orange", "grapes")                 //append is used to add a new element to the slice
	fmt.Printf("fruitlist is %T\n ", fruitlist)
	fmt.Println("fruitlist is ", fruitlist)
	fmt.Println("fruitlist length is ", len(fruitlist))

	fruitlist = append(fruitlist[1:]) //this is used to remove the first element of the slice
	fmt.Println("fruitlist after removing first element is ", fruitlist)

	fruitlist = append(fruitlist[:3]) //0,1,2 taken indexes
	fmt.Println("fruitlist shortened to 3 elements is ", fruitlist)

	highscore := make([]int, 4) //make is used to create a slice with a specific length
	highscore[0] = 100
	highscore[1] = 200
	highscore[2] = 300
	highscore[3] = 400
	fmt.Println("highscore is ", highscore)
	highscore = append(highscore, 500, 600) //adding more elements to the highscore slice
	fmt.Println("highscore after adding more elements is ", highscore)

	sort.Ints(highscore)                                  //sort is used to sort the slice in ascending order
	fmt.Println("highscore after sorting is ", highscore) //this will only available in the slice package

	fmt.Println(sort.IntsAreSorted(highscore)) //check if the slice is sorted or not
}
