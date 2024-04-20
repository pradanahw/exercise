package employee

// Staff is a struct for staff
// It embeds Employee
// This means that Staff has all the fields and methods of Employee
type Staff struct {
	Employee // Embedding Employee struct
}

// NewStaff creates a new staff
// It returns a pointer to the staff
// Creational method
func NewStaff() *Staff {
	// Create a new staff
	staff := &Staff{}

	// Set the name to "Staff"
	staff.Name = "Staff"

	// Set the salary to 500
	staff.Salary = 500

	return staff
}
