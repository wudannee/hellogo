package ch4

import (
	"fmt"
	"time"
)

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

	var personArr2 [2]Person

	// the following line will cause compile error
	// cannot assign [5]Person to [2]Person
	// personArr2 = personArr

	// copy the elements of personArr to personArr2
	copy(personArr2[:], personArr[:])
	fmt.Println("personArr2:", personArr2)

	// two-dimensional array
	var twoDimArr = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for _, v := range twoDimArr {
		for _, v2 := range v {
			fmt.Printf("%d ", v2)
		}
		fmt.Println()
	}

	var row2 [3]int = twoDimArr[1]
	fmt.Println("row2 of twoDimArr:", row2)

	line := ""
	fmt.Scanf("%s", &line)

	// largeArr takes up a lot of memory
	var largeArr [N]int
	ExpensiveArrayCopy(largeArr)

	fmt.Println("exit after 5 seconds")
	time.Sleep(time.Second * 5)
}

// stack overflow
// const N = 1024 * 1024 * 100
const N = 1024 * 1024 * 10

func ExpensiveArrayCopy(arr [N]int) {
	fmt.Println("in a function, arr:", len(arr))
	for i := 0; i < len(arr); i++ {
		arr[i] += 1
	}
	for {
		line := ""
		fmt.Scanf("%s", &line)
		if line == "exit" {
			break
		}
	}
}
