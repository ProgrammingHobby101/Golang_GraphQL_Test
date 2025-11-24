package main

import "fmt"

func changeX1(y *int) {
	*y = 35 //change to 35
	fmt.Println("Value of y in ChangeX1:", *y)
}
func changeX2(a *int) {
	*a = 55
	fmt.Println("Value of a in ChangeX2:", *a)
}

func main() {
	x := 10
	ptr := &x // ptr holds the address of x
	fmt.Println("Value of x:", x)
	fmt.Println("Value pointed to by ptr (*ptr):", *ptr) // Dereferencing ptr to get the value of x
	fmt.Println("ptr address: ", &ptr)
	*ptr = 20 // Modify the value at the address ptr points to
	fmt.Println("New value of x:", x)
	changeX1(ptr) //ptr here must be a pointer before sending it to changeX1.
	fmt.Println("Second to last value of x:", x)
	z := x       // pass by value.
	changeX2(&z) //z must be a pointer here before sending it to changeX2, this means that *z  won't work in this function call.
	fmt.Println("Final value of x in after ChangeX2:", x)//35
	fmt.Println("Final value of z in after ChangeX2:", z)//55
}
