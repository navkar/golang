package main
import (
  "fmt"
)

type Employee struct {
  FirstName string
  LastName string
  Role string
}

//the type *Employee is the receiver of the ManagerRole method.
func (e *Employee) ManagerRole() {
  e.Role = "manager"
}

func (e *Employee) EmployeeRole() {
  e.Role = "employee"
}

// factory function to create a new instance
func NewEmployee(firstName string, lastName string) *Employee {
  return &Employee {
    FirstName: firstName,
    LastName: lastName,
    Role: "" }
}

//goku := new(Saiyan) // same as
//goku := &Saiyan{}

func main() {
  mgr := NewEmployee ("Bill", "Gates")
  mgr.ManagerRole()

  emp := new (Employee)
  emp.FirstName = "Satya"
  emp.LastName = "Nadella"
  emp.EmployeeRole()

  fmt.Println("FirstName | LastName | Role")
  fmt.Println(mgr)
  fmt.Println(emp)
}
