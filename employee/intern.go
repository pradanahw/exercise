package employee

// Intern is a struct for intern
// It embeds Employee
// This means that Intern has all the fields and methods of Employee
type Intern struct {
	Employee // Embedding Employee struct
}

// NewIntern creates a new intern
// It returns a pointer to the intern
// Creational method
func NewIntern() *Intern {
	// Create a new intern
	intern := &Intern{}

	// Set the name to "Intern"
	intern.Name = "Intern"

	// Set the salary to 100
	intern.Salary = 100

	return intern
}
