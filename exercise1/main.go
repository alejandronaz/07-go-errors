package main

import (
	"clase7/salaryError"
	"fmt"
)

func main() {

	salary := 10000
	err := CalculateTax(salary)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func CalculateTax(salary int) error {
	if salary < 150000 {
		return salaryError.SalaryError{Code: 1}
	}
	fmt.Println("Must pay tax")
	return nil
}
