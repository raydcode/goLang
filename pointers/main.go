package main

import "fmt"

func main() {
 
	 a := 5
	 b := &a // Memory address of variable a 

	fmt.Println(a,b)

	fmt.Printf("%T\n",b) // * represents pointer

	// Use * to read values from memory address

	fmt.Println(*b)

	// Change value to pointer

	*b=10

	fmt.Println(a)
}
