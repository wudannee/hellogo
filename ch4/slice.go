package ch4

import "fmt"

func DemoSlice() {
	slice := make([]int, 3)
	slice[0] = 1
	fmt.Println("slice:", slice)
}
