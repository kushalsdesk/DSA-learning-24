package main

import "fmt"

type Car struct {
	Maker string
	Model string
	Year  int
	Color string
}

func (c *Car) showDetails() {
	fmt.Printf("The Car Model is %s from %s from the year %d with color %s\n", c.Model, c.Maker, c.Year, c.Color)
}

// a new practice begins with banking system
type Person struct {
	Name string
	Age  int
}
type Account struct {
	Owner     Person
	AccNumber int
	AccValue  float64
}

func (acc *Account) deposit() float64 {
	var newValue float64
	fmt.Println("Enter the Amount to deposit: ")
	fmt.Scan(&newValue)
	acc.AccValue += newValue
	return acc.AccValue
}
func (acc *Account) withdraw() float64 {
	var valueToDebit float64
	fmt.Println("Enter the Amount to withdraw: ")
	fmt.Scan(&valueToDebit)

	if valueToDebit > acc.AccValue {
		fmt.Println("Your Account Doesnot Have enough to withdraw")
	} else {

		acc.AccValue -= valueToDebit
	}
	return acc.AccValue
}

func main() {
	//Code for Car structure
	lambo := Car{Maker: "Lamborghini", Model: "Countach", Year: 2005, Color: "Yellow"}
	lambo.showDetails()

	// code for banking system

	acc := Account{
		Owner:     Person{Name: "Kushal", Age: 24},
		AccNumber: 2098,
		AccValue:  2020.090,
	}
	fmt.Printf("After Withdrawl new balance: %f\n", acc.withdraw())
	fmt.Printf("After Deposit new balance: %f\n", acc.deposit())

}
