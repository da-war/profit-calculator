package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Please Enter your revenue:")
	fmt.Scan(&revenue)

	fmt.Print("Please Enter your expenses:")
	fmt.Scan(&expenses)

	fmt.Print("Please Enter your tax rate:")
	fmt.Scan(&taxRate)

	earningBeforeTax := revenue - expenses
	profit := earningBeforeTax * (1 - taxRate/100)
	ratio := earningBeforeTax / profit

	fmt.Println("Earning Before Tax:", earningBeforeTax)
	fmt.Println("Profit", profit)
	fmt.Println("Ratio", ratio)

}
