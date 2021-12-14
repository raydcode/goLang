package main

import "fmt"

/*
. A closure is a special type of anonymous function that references variables declared outside of the function itself. It is similar to accessing global variables which are available before the declaration of the function.

	// Declaring the variable
	GFG := 0

	// Assigning an anonymous
	// function to a variable
	counter := func() int {
	GFG += 1
	return GFG
	}

	fmt.Println(counter())
	fmt.Println(counter())
	
	
}


*/

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func main() {

	sum := adder()

	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}

}
