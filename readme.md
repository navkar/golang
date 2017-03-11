## Welcome

I have written some code snippets while learning golang. Look around and learn!

### main function in golang

Every Go program starts in a package main by calling identically named function. When this function returns the program ends its execution.

```go
  package main
  import "fmt"

  func main() {
      fmt.Printf("Hello from the GO programming language\n")
      fmt.Printf("This is lot like C programming language\n")
  }
```

### golang commands on bash

```bash
go run hello-world.go
```

```bash
go build hello-world.go
./hello-world
```

### init function in golang and more on packages

* init functions are defined in package block.
* init function doesn’t take arguments neither returns any value.
* identifier init is not declared so cannot be referenced.
* Many init functions can be defined in the same package.
* init identifier doesn’t introduce binding, just as blank identifier represented by underscore character.
* Package initialization is done only once even if package is imported many times.
* To use a imported package it needs to be initialized first.
* Package in Go can contain many files. To ensure reproducible initialization behavior, build systems are encouraged to present multiple files belonging to the same package in lexical file name order to a compiler.
* init is called after all the global variables in the package (all files) are initialized.
* The most common use case of init function is to assign a value which cannot be calculated as a part of initialization expression
* for instance, it’s not possible to use for loop as an expression (it’s a statement in Go) so putting it into init function solves the problem.

```go
  var precomputed = [20]float64{}
  func init() {
      var current float64 = 1
      precomputed[0] = current
      for i := 1; i < len(precomputed); i++ {
          precomputed[i] = precomputed[i-1] * 1.2
      }
  }
```

### solving unused imports with a blank identifier

* Go is very strict when it comes to unused variables/imports.

```go
  import _ "ThisIsGoingToBuild"
```

### What is the difference between the following code snippets?

#### snippet #1

```go
    package main
    import "fmt"
    var (
        a int = b + 1
        b int = 1
    )
    func main() {
        fmt.Println(a)
        fmt.Println(b)
    }
```

#### snippet #2

```go
    package main
    import "fmt"
    func main() {
        var (
            a int = b + 1
            b int = 1
        )
        fmt.Println(a)
        fmt.Println(b)
    }
```


### golang commands on bash

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
* [go-lang-cheat-sheet] (https://github.com/a8m/go-lang-cheat-sheet/)

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
