package main

import (
"fmt"
)

// Functions on Structures
type Employee struct {
  FirstName string
  LastName string
  Name string
}

// The type *Employee is the receiver of the FullName method
func (e *Employee) FullName() {
   e.Name = e.FirstName + " " + e.LastName
}

func main(){
  emp := Employee {"naveen", "karam", "" }
  emp.FullName()

  fmt.Println("full name: " + emp.Name)
}
