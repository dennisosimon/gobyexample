package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	// checks if key exists in map, ignores the value with _
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// that's how we get the value, 0 because not present
	prs1, prs2 := m["k2"]
	fmt.Println("prs1:", prs1)
	fmt.Println("prs2:", prs2)

	n := map[string]int {"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
