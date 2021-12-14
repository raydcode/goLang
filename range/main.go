package main

import "fmt"

func main() {
	ids := []int{1, 2, 3, 4, 5}


	// Loop through ids using range

	for i ,id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}
 
	// Not using Index

	for _ ,id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add all the ID'S

	sum := 0

	for _ ,id := range ids {
		sum += id
	}

	fmt.Println("Total of ID",sum)

	// range with map

	emails := map[string]string{
		"kathir":"kathir@gmail.com",
		"karthik":"karthik@gmail.com",
	}

	for k ,v := range emails {
		fmt.Printf("%s: %s\n",k,v)
	}

	for k := range emails {
		fmt.Println("name: "+k)
	}
}
