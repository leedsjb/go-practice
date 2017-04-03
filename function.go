package main

import "fmt"

func add(x int, y int) (int, int) {
	fmt.Println("executing add function")
	return 0, x + y
}
