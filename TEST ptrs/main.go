package main

import "fmt"

func main() {
	x := 10
	ptr := &x // ptr holds the address of x
	fmt.Println("Value of x:", x)
	fmt.Println("Value pointed to by ptr (*ptr):", *ptr) // Dereferencing ptr to get the value of x
	fmt.Println("ptr address: ", &ptr)
	*ptr = 20 // Modify the value at the address ptr points to
	fmt.Println("New value of x:", x)

}
