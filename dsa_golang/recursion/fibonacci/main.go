package main

import "fmt"

var target int

func main() {
	fmt.Println("Enter the target to calculate series to that: ")
	fmt.Scan(&target)
	fmt.Printf("The series to %d is: ", target)
	for i := 1; i <= target; i++ {
		fmt.Printf(" %d", fibo(i))
	}
	fmt.Println()
}

func fibo(target int) int {
	if target == 1 || target == 2 {
		return 1
	}
	return fibo(target-1) + fibo(target-2)
}
