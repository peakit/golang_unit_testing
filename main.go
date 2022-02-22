package main

import (
	"fmt"
)

func main() {
	p1 := Initialize(2, 3)
	p2 := Initialize(3, 4)
	fmt.Printf("First point: %+v\n", p1)
	fmt.Printf("Second point: %+v\n\n", p2)

	p3 := Initialize(7, 3)
	p4 := Initialize(5, 2)
	fmt.Printf("Third point: %+v\n", p3)
	fmt.Printf("Fourth point: %+v\n\n", p4)
}
