package main

import "fmt"

func main() {

   
	 var userArr [2]string

	 userArr[0] = "kathir"
	 userArr[1] = "karthik"
     
	 // Declaration and assing 
	 nameArr := [2]string{"kathir", "karthik"}

	 

     fmt.Println(userArr)

	 fmt.Println(nameArr)

   //Slice
   movieArr := []string{"Liar Liar", "School of Rock","Young Sheldon"}
   fmt.Println(movieArr)

   // use len() to get the length of Array

   fmt.Println(len(movieArr))

   // use range  Array[start:end] of index in a array

   fmt.Println(movieArr[1:3])

}
