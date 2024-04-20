package employee

// Manager is a struct for manager
// It embeds Employee
// This means that Manager has all the fields and methods of Employee
type Manager struct {
	Employee // Embedding Employee struct
}

// NewManager creates a new manager
// It returns a pointer to the manager
// Creational method
func NewManager() *Manager {
	// Create a new manager
	manager := &Manager{}

	// Set the name to "Manager"
	manager.Name = "Manager"

	// Set the salary to 1000
	manager.Salary = 1000

	return manager
}
