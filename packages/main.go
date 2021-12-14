package main

import (
	"fmt"
	"math"
	"github.com/raydcode/packages/strutil"
)

// use absolute paths for packages


func main() {
	//** pretty common Math functions in all languages :

	fmt.Println(math.Floor(2.5))
	fmt.Println(math.Ceil(2.5))
	fmt.Println(math.Sqrt(16))

	// Custom packages
    fmt.Println(strutil.Reverse("olleh"))
   
   

}
