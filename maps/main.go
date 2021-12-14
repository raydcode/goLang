package main

import "fmt"

func main() {
	// Define maps

	// email := make(map[string]string)

	//Assign key  values

	// email["kathir"] = "kathir@gmail.com"
	// email["karthik"] = "karthik@gmail.com"


	// Declare & Assign 

	email := map[string]string{
		"kathir":"kathir@gmail.com",
		"karthik":"karthik@gmail.com",
	}

	fmt.Println(email)

	fmt.Println(email["kathir"])
	// you also use len() like in arary
	fmt.Println(len(email))

	// delete map
	delete(email, "karthik")

	fmt.Println(email)

}
