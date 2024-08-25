package main

import "fmt"

func main() {
	var number, sum int
	fmt.Println("Enter the Number for sum of digits: ")
	fmt.Scan(&number)
	// geting another variable to keep the original intact
	clone_num := number

	for clone_num > 0 {
		sum += clone_num % 10
		clone_num = clone_num / 10
	}

	fmt.Printf("The sum of digits for %d is %d\n", number, sum)
}
