package main

import (
	"clase7/salaryError"
	"errors"
	"fmt"
)

func main() {

	salary := 8000

	err := CalculateTax(salary)
	errKnown := salaryError.SalaryError{Code: 2}

	if errors.Is(err, errKnown) {
		fmt.Println(err)
		return
	}

}

func CalculateTax(salary int) error {
	if salary < 10000 {
		return salaryError.SalaryError{Code: 2}
	}
	fmt.Println("Must pay tax")
	return nil
}
