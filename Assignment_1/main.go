package main

import "fmt"

func main() {
	array1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for index, values := range array1 {

		if array1[index]%2 == 0 {
			fmt.Println(values, "is even")
		} else if array1[index]%2 != 0 {
			fmt.Println(values, "is odd")
		} else {
			fmt.Println("invalid input")
		}
	}
}
