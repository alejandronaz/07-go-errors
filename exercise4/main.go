package main

import (
	"errors"
	"fmt"
)

var errSalary = errors.New("salary error")

func main() {

	salary := 8000

	err := CalculateTax(salary)

	if errors.Is(err, errSalary) {
		fmt.Println(err)
		return
	}

}

func CalculateTax(salary int) error {
	if salary < 150000 {
		return fmt.Errorf("%w: the minimum taxable amount is 150,000 and the salary entered is: %d", errSalary, salary)
	}
	fmt.Println("Must pay tax")
	return nil
}
