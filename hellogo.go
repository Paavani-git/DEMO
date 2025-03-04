package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("Hello\n") //no nextline
	fmt.Println("Hello") // with nextline
	fmt.Printf("%T\n", 23.2)
	fmt.Printf("%v\n", []int{1, 2, 3})  //default format
	fmt.Printf("%#v\n", []int{1, 2, 3}) //debugging mode
	fmt.Fprintf(os.Stderr, "This is a formatted error: %d\n", 404)
}

/*
%v - Default format
%+v - Detailed struct with field names
%#v - Go-syntax representation
%T - Type of the value
%d - Decimal integer
%f - Floating point
%s - String
%q - Double-quoted string
%p - Pointer address
*/
