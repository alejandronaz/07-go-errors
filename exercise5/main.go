package main

import (
	"clase7/salaryError"
	"fmt"
)

func calculateSalary(horas int, valorHora float64) (salary float64, err error) {

	if horas < 80 {
		return 0, salaryError.SalaryError{Code: 3}
	}

	salary = float64(horas) * valorHora
	if salary >= 150000 {
		salary *= 0.9
	}
	return salary, nil

}

func main() {

	fmt.Println("---- case 1 ----")
	salary, err := calculateSalary(100, 20000)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Salario: ", salary)

	fmt.Println("---- case 2 ----")
	salary2, err2 := calculateSalary(60, 20000)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println("Salario: ", salary2)
}
