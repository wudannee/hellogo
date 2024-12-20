package ch4

import "fmt"

func is_nilmap(m map[string]int) bool {
	if m == nil {
		return true
	}
	return false
}
func DemoMap() {
	m := map[string]int{"a": 1, "b": 2}
	fmt.Println(m)

	var m2 map[string]int
	fmt.Printf("m2 is nil? %v,  %p\n", is_nilmap(m2), m2)

	m3 := make(map[string]int)

	fmt.Printf("m3 is nil? %v,  %p\n", is_nilmap(m3), m3)
}
