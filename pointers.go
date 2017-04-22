package main

import (
	"fmt"
	"strconv"
)

func pointers() {

	i := 42
	// j := 2701

	var p *int = &i // p points to the memory address of i
	fmt.Println(i)
	fmt.Println(p)  // memory address e.g.: 0xc420076178
	fmt.Println(*p) // 42

	*p = 21
	fmt.Println("i is: " + strconv.Itoa(i)) // 21
	fmt.Printf("p is: %v \n", p)            // 0xc420076178
	fmt.Println("*p is: " + strconv.Itoa(*p))

	// p = j
	// fmt.Println(p)

	// p /= 37
	// fmt.Println(j)
}
