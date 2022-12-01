package main

import (
	"fmt"
)

func main() {

	/*	m := new(map[string]int)
			m["theAnswer"] = 42
			fmt.Println(m)


		m := make(map[string]int)
		m["theAnswer"] = 42
		fmt.Println(m)
	*/
	fmt.Println("Pointers")

	var pp *int                      // * means pointer, not an int
	fmt.Println("Value of pp: ", pp) //Will be nil as

	anInt := 42
	var p = &anInt // & points at the memmory address of the variable, not its value
	fmt.Println("Value of p:", *p)

	value1 := 42.13
	pointer1 := &value1
	fmt.Println("Value 1: ", *pointer1)

	*pointer1 = *pointer1 / 31 //updating the value of the variable, using the pointer.
	//value of variable at location x = value of variable at location x / 31

	fmt.Println("Pointer 1:", *pointer1)
	fmt.Println("Value 1:", value1)
}
