package main

type Employee struct {
	Name  string
	Age   int
	Email string
}

// Employee builder pattern code
type EmployeeBuilder struct {
	employee *Employee
}

func NewEmployeeBuilder() *EmployeeBuilder {
	employee := &Employee{}
	b := &EmployeeBuilder{employee: employee}
	return b
}

func (b *EmployeeBuilder) Name(name string) *EmployeeBuilder {
	b.employee.Name = name
	return b
}

func (b *EmployeeBuilder) Age(age int) *EmployeeBuilder {
	b.employee.Age = age
	return b
}

func (b *EmployeeBuilder) Email(email string) *EmployeeBuilder {
	b.employee.Email = email
	return b
}

func (b *EmployeeBuilder) Build() (*Employee, error) {
	return b.employee, nil
}
