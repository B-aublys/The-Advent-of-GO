package main

import "fmt"

// The rule about pointers vs. values for receivers is that
// value methods can be invoked on pointers and values, but
// pointer methods can only be invoked on pointers. This is
// because pointer methods can modify the receiver; invoking
// them on a copy of the value would cause those modifications
// to be discarded.

type structure struct {
	x int
}

func (st structure) change_x(new_x int) {
	st.x = new_x
}

func (st structure) String() string {
	return "hello mister"
}

type structure2 struct {
	x int
}

func (st *structure2) change_x(new_x int) {
	st.x = new_x
}

func (st structure2) String() string {
	return "hello mister"
}

func main() {
	x := structure{5}
	x_p := &structure{5}

	x2 := structure2{5}
	x2_p := &structure2{5}

	x.change_x(1)
	x_p.change_x(1)

	x2.change_x(1)
	x2_p.change_x(1)

	fmt.Printf("x: %d,  x_p: %d\n", x.x, x_p.x)
	fmt.Printf("x2: %d,  x2_p: %d\n", x2.x, x2_p.x)

	fmt.Println(x2_p)
}
