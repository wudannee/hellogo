package ch4

import "fmt"

type Person struct {
	Name string
	Age  int
}

func DemoArray() {
	intArr := [3]int{1, 2, 3}
	intArr2 := [...]int{3, 2, 1}
	fmt.Println(intArr, intArr2)

	var floatArr = [...]float64{1.1, 2.2, 3.3}
	var total float64
	for _, v := range floatArr {
		total += v
	}
	fmt.Println("total:", total)

	var personArr = [5]Person{
		{Name: "Suan", Age: 16},
		{Name: "John", Age: 18},
	}

	// test the copy of an array
	var personArrCopy [5]Person

	personArrCopy = personArr
	// modify the original array
	personArr[0].Name = "Andy"
	// modify the copied array
	personArrCopy[0].Name = "Bob"

	// print the original array and the copied array
	fmt.Println("modified personArr:", personArr)
	fmt.Println("modified personArrCopy:", personArrCopy)

}

