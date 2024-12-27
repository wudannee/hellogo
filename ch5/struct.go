package ch5

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func DemoStruct() {
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

	// create a struct variable
	p := Person{Name: "Susan", Age: 16}
	fmt.Println("p:", p)
	// When a method with a pointer receiver is called on a value, the compiler automatically takes the address of the value and passes a pointer to the method.
	p.SetName("John")
	fmt.Println("john?", p)

	// create a pointer to the struct variable
	fmt.Println("create a pointer to the struct variable")
	pr := &p
	pr.SetName("Lucy")
	fmt.Println("after change p:", pr)
}
