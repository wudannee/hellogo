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

	// var p2 = new(Person)
	// try to call a method on a nil pointer to struct, and fix the panic issue
	var p2 *Person
	operateP2 := func() {
		p2.SetName("Tom")
		fmt.Println("p2:", p2)
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("ERROR:", err)
			// possibly p2 is nil, so we need to new it before calling operateP2
			p2 = new(Person)
			operateP2()
		}

		// resume the demo when the panic is handled
		ResumeDemoStruct()
	}()
	operateP2()
}

func ResumeDemoStruct() {
	fmt.Println("resume demo struct")
}
