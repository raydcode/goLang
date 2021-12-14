package main

import (
	"fmt"
	
)

func main() {
	x := 15
	y := 10

	//** use can able to paranthesis (No err throws)

	if x < y {
		fmt.Printf("%d is less than  %d\n", x, y)
	} else {
		fmt.Printf("%d is greater than  %d\n", x, y)
	}

	color := "red"

	if color == "red"{
		fmt.Println("color is red")
	}else if color == "green"{
		fmt.Println("color is green")
	}else{
		fmt.Printf("color is not red or green")
	}

    switch color {
		case "green":
			fmt.Println("color is green")
		case "yellow":
			fmt.Println("color is green")
		default:
			fmt.Println("color is green or yellow")	
	} 

}
