package employee

// Director is a struct for director
type Director struct {
	Employee
}

// NewDirector creates a new director
func NewDirector() *Director {
	director := &Director{}
	director.Name = "Director"
	director.Salary = 5000
	return director
}

// CalculateBonus calculates bonus for Director
func (d *Director) CalculateBonus() {
	d.Bonus = float64(d.Salary) * 0.30 // 30% dari gaji
}
