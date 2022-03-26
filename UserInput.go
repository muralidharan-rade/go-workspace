package main

import (
"fmt";
"reflect"	
)

func main() {
	
	var name string
	fmt.Printf("Enter your name: ")
	fmt.Scanf("%v", &name)
	
	fmt.Println("User entered name: ", name)
	
	// TypeOf
	fmt.Println(reflect.TypeOf(name))
	fmt.Println(reflect.TypeOf(20.0))
}

