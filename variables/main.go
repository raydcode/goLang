package main

import "fmt"



func main() {
   // Main Types
   //string
   //bool
   //int int8 int16 int32 int64 
   //uint uint8 uint16 uint32 uint64 uintptr
   //byte - alias for uint8
   //rune - alias for  int32
   //float32 float64
   //complex64 complex128

   //Using variables



   var name string = "Kathir"
   var age int = 23;
   const IsActive = true;

    // ** works fine with  without explict datatype
	// Like Javascript you can use constant variables


	//Shorthand for variables

	role := "developer"
	ctc:=3.6
	//** Shorthand variables only used inside functions (can't global)
  
	// continous variables declaration
    gender,place := "male", "ramanathapuram"


  
   fmt.Println(name,age,IsActive,role,ctc,gender,place);
   fmt.Printf("%T\n",gender);
}
