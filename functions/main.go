package main

import "fmt"


func sayHello(name string) string {

	return "Hello " + name
	
}

func getSum(num1 int , num2 int)  int{
	return num1 + num2
}

func main() {
	fmt.Println(sayHello("kathir"));
	fmt.Println(getSum(3,3))
}
