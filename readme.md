## GO commands on bash

```bash
go run hello-world.go
go build hello-world.go
./hello-world
```

### GO Lang references
* [multiple-return-values] (https://gobyexample.com/multiple-return-values)
* [golang] (https://golang.org/)
* [overview-of-Go] (https://www.codeproject.com/Articles/1160635/Quick-Overview-of-Go)
* [golang-book] (https://www.golang-book.com/books/intro/10)


### Struct in go - Employee-Manager example
```go
package main
import (
  "fmt"
)

type Employee struct {
  FirstName string
  LastName string
  Role string
  Manager *Employee
}

func (e *Employee) SetManager(mgr *Employee) {
  e.Manager = mgr
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
    Role: "",
    Manager: nil }

    // emp := new (Employee)
    // emp.FirstName = firstName
    // emp.LastName = lastName
    // return emp
}

//emp := new (Employee) // same as
//emp := &Employee{}

func main() {
  mgr := NewEmployee ("Bill", "Gates")
  mgr.ManagerRole()

  emp := new (Employee)
  emp.FirstName = "Satya"
  emp.LastName = "Nadella"
  emp.EmployeeRole()
  emp.SetManager(mgr)

  fmt.Println("FirstName | LastName | Role | Manager")
  fmt.Println(mgr)
  fmt.Println(emp)
  fmt.Println(emp.Manager.FirstName + " is the manager of " + emp.FirstName)
}
```

#### Output 

```bash
Naveens-MacBook-Pro:structs navkar$ go run employee.go
FirstName | LastName | Role | Manager
&{Bill Gates manager <nil>}
&{Satya Nadella employee 0xc420014200}
Bill is the manager of Satya
Naveens-MacBook-Pro:structs navkar$
```
