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

### what are blocks in golang?

* Block is a sequence of statements (empty sequence is also valid). Blocks can be nested and are denoted by curly brackets

```go
  package main
  import “fmt”
  func main() {
      { // start outer block
          a := 1
          fmt.Println(a)
          { // start inner block
              b := 2
              fmt.Println(b)
          } // end inner block
      } // end outer block
  }
```

* universe block contains all source code
* package block contains all package’s source code (package can be spread over several files in a single directory)
* file block contains file’s source code (importing packages names have scope set to file block)
* for statement is in its own implicit block

```go
  for i := 0; i < 5; i++ {
      fmt.Println(i)
  }
```

```go
  if i := 0; i >= 0 {
      fmt.Println(i)
  }
```

```go
  switch i := 2; i * 4 {
  case 8:
      fmt.Println(i)
  default:
      fmt.Println(“default”)
  }
```

```go
  v := 1
  {
      v := 2  // short variable declaration
      fmt.Println(v)
  }
  fmt.Println(v)
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

### Struct in golang - Employee-Manager example
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
### golang interfaces

* an interface is a set of methods that we use to define a set of actions
* an interface defines a set of methods, so if a type implements all the methods we say that it implements the interface.
* an interface can be implemented by any type, and one type can implement many interfaces simultaneously.
* any type implements the empty interface `interface{}` because it doesn't have any methods and all types have zero methods by default.

```go
  // define a as empty interface
  var a interface{}
  var i int = 5
  s := "Hello world"
  // a can store value of any type
  a = i
  a = s
```

* an interface is a set of abstract methods, and can be implemented by non-interface types.
* an empty interface is an interface that doesn't contain any methods, so all types implement an empty interface. This fact is very useful when we want to store all types at some point, and is similar to `void*` in C.
* If a function uses an empty interface as its argument type, it can accept any type; if a function uses empty interface as its return value type, it can return any type.

### type of a variable in an interface

If a variable is the type that implements an interface, we know that any other type that implements the same interface can be assigned to this variable. The question is how can we know the specific type stored in the interface.

#### Assertion of Comma-ok pattern

Go has the syntax `value, ok := element.(T)`. This checks to see if the variable is the type that we expect, where "value" is the value of the variable, "ok" is a variable of boolean type, "element" is the interface variable and the T is the type of assertion.
If the element is the type that we expect, ok will be true, false otherwise.

```go
  package main

  import (
      "fmt"
      "strconv"
  )

  type Element interface{}
  type List []Element

  type Person struct {
      name string
      age  int
  }

  func (p Person) String() string {
      return "(name: " + p.name + " - age:     " + strconv.Itoa(p.age) + " years)"
  }

  func main() {
      list := make(List, 3)
      list[0] = 1       // an int
      list[1] = "Hello" // a string
      list[2] = Person{"Dennis", 70}

      for index, element := range list {
          if value, ok := element.(int); ok {
              fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
          } else if value, ok := element.(string); ok {
              fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
          } else if value, ok := element.(Person); ok {
              fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
          } else {
              fmt.Printf("list[%d] is of a different type\n", index)
          }
      }
  }
```

#### Assertion of switch pattern

* `element.(type)` cannot be used outside of the `switch` body, which means in that case you have to use the comma-ok pattern .

```go
  func main() {
      list := make(List, 3)
      list[0] = 1       //an int
      list[1] = "Hello" //a string
      list[2] = Person{"Dennis", 70}

      for index, element := range list {
          switch value := element.(type) {
          case int:
              fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
          case string:
              fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
          case Person:
              fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
          default:
              fmt.Println("list[%d] is of a different type", index)
          }
      }
  }
```

### Embedded Interfaces in golang


### golang references
* [multiple-return-values] (https://gobyexample.com/multiple-return-values)
* [golang] (https://golang.org/)
* [overview-of-Go] (https://www.codeproject.com/Articles/1160635/Quick-Overview-of-Go)
* [golang-book] (https://www.golang-book.com/books/intro/10)
* [go-lang-cheat-sheet] (https://github.com/a8m/go-lang-cheat-sheet/)
