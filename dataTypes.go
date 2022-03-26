package main

import (
"fmt";
"strconv"
)

const PI float64 = 3.14

func main() {
	
	fmt.Println(PI)
	
	fmt.Printf("%f", PI)
	fmt.Println("\n")

	// type casting	
	var f float64 = 3.14
	
	i := int(f)
	fmt.Println(i)
	
	
	// strconv package
	s := "200"
	k, err := strconv.Atoi(s)
	fmt.Printf("string to integer : %d %T", k, k)
	fmt.Printf("string to integer.. any errors? 	: %v", err)
	fmt.Println("\n")
	
	// different data types with default values
	
	var ii int
	var ff float64
	var ss string
	var bb bool
	fmt.Printf("different types :: %d %T \n", ii, ii)
	fmt.Printf("different types :: %f %T \n", ff, ff)
	fmt.Printf("different types :: %v %T \n", ss, ss)
	fmt.Printf("different types :: %t %T \n", bb, bb)
	
	
	// values assigned
	ii = 1000
	ff = 10.0
	ss = "Hello World"
	bb = true
	
	fmt.Printf("different types :: %d %T \n", ii, ii)
	fmt.Printf("different types :: %f %T \n", ff, ff)
	fmt.Printf("different types :: %v %T \n", ss, ss)
	fmt.Printf("different types :: %t %T \n", bb, bb)
}
