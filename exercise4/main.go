package main

import (
	"fmt"
)

func main() {

	salary := 8000

	err := CalculateTax(salary)

	if err != nil {
		fmt.Println(err)
		return
	}

}

func CalculateTax(salary int) error {
	if salary < 150000 {
		return fmt.Errorf("error: the minimum taxable amount is 150,000 and the salary entered is: %d", salary)
	}
	fmt.Println("Must pay tax")
	return nil
}
