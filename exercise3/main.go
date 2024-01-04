package main

import (
	"errors"
	"fmt"
)

var errSalary = errors.New("error: salary is less than 10000")

func main() {

	salary := 8000

	err := CalculateTax(salary)

	if errors.Is(err, errSalary) {
		fmt.Println(err)
		return
	}

}

func CalculateTax(salary int) error {
	if salary < 10000 {
		return errSalary
	}
	fmt.Println("Must pay tax")
	return nil
}
