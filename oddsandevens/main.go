package main

import "fmt"

func main() {
	var oneToTen []int
	for i := 1; i <= 10; i++ {
		oneToTen = append(oneToTen, i)
	}

	for _, num := range oneToTen {
		if num%2 == 0 {
			fmt.Printf("%v is even", num)
		} else {
			fmt.Printf("%v is odd", num)
		}
	}
}
