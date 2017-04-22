package main

import "fmt"

//Vertex a struct to hold the two coordinates that define a vertex
type Vertex struct {
	X int
	Y int
}

// UpdateX function to update the X field of the struct upon which its called
func (vtx *Vertex) UpdateX(x int) {
	vtx.X = x
}

func structs() {

	p := Vertex{ // longhand notation to create a new struct
		X: 10,
		Y: 10,
	}

	r := &p // create a pointer to struct p at address &p

	p.X = 15 // can access struct fields using dot notation

	(*r).Y = 25
	r.Y = 25 // can omit the * from *r when accessing a struct field without the explicit dereference(*)

	fmt.Println(p.X)
	p.UpdateX(20)
	fmt.Println(p.X)

	fmt.Println(p.Y)

	q := Vertex{30, 40} // shorthand notation to create a new struct w/o explicitly mentioning field names
	fmt.Println(q)
}
