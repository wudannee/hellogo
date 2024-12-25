package ch5

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func DemoStruct() {
	p := Person{Name: "Susan", Age: 16}
	fmt.Println("p:", p)

	var k []int
	if k == nil {
		fmt.Println("k is nil", k)

		fmt.Println("append 10 to k")
		k = append(k, 10)
		fmt.Println("k is not nil now", k, len(k), cap(k))
	}

	j := make([]int, 0, 5)
	if j != nil {
		fmt.Println("j is not nil", j, len(j), cap(j))
	}
}
