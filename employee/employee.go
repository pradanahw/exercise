package employee

// EmployeeInterface is an interface for employee
type EmployeeInterface interface {
	SetSalary(salary int)
	GetName() string
	GetSalary() int
	GetBonus() float64
}

// Employee is a struct for employee
type Employee struct {
	Name   string
	Salary int
}

// SetSalary sets the salary of the employee
func (e *Employee) SetSalary(salary int) {
	e.Salary = salary
}

// GetName gets the name of the employee
func (e *Employee) GetName() string {
	return e.Name
}

// GetSalary gets the salary of the employee
func (e *Employee) GetSalary() int {
	return e.Salary
}

// GetBonus gets the bonus of the employee
func (e *Employee) GetBonus() float64 {
	// Implement logic to calculate bonus
	// For example, bonus is calculated as 10% of salary
	bonusPercentage := 0.10
	return float64(e.Salary) * bonusPercentage
}
