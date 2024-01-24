// Sample program to show how unexported fields from an exported
// struct type can't be accessed directly.
package main

import (
	"fmt"

	"github.com/goinaction/code/chapter5/listing74/entities"
)

// main is the entry point for the application.
func main() {
	// Create a value of type Admin from the entities package.
	a := entities.Admin{
		Rights: 10,
	}

	// Set the exported fields from the unexported
	// inner type. 将内部类型提升到外部，由于Name和Email声明是公开的所以可以直接访问
	a.Name = "Bill"
	a.Email = "bill@email.com"

	fmt.Printf("User: %v\n", a)
}
