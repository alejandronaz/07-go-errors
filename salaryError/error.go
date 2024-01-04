package salaryError

type SalaryError struct {
	// message string
	Code int
}

func (e SalaryError) Error() string {
	switch e.Code {
	case 1:
		return "Error: the salary entered does not reach the taxable minimum"
	case 2:
		return "Error: salary is less than 10000"
	case 3:
		return "Error: the worker cannot have worked less than 80 hours per month"
	default:
		return "Error: unknown"
	}
}
