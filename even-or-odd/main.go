package main

import (
	"fmt"
)

func main(){
	numbers := []int {1,2,3,4,5,6,7,8,9,10}

	for _,number := range numbers {

		if number % 2 == 0 {
			fmt.Println(number, "is an even number")
			continue
		}

		fmt.Println(number,"is an odd number")
	}
}