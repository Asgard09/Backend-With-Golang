package main

import "fmt"
type Vertex struct {
	X int
	Y int
	Z int
}

var m map[string][]int

func main() {
	m = make(map[string][]int)
	m["google"] = []int{1, 2, 3}
	m["microsoft"] = []int{4, 5, 6}
	fmt.Println(m)

	// delete an entry from the map
	delete(m, "google")
	fmt.Println(m)

	// check if a key exists in the map
	v, ok := m["microsoft"]
	fmt.Println(v, ok)

}