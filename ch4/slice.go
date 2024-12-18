package ch4

import "fmt"

func DemoSlice() {
	slice := make([]int, 3)
	slice[0] = 1
	fmt.Println("slice:", slice)

	// declare a nil slice
	var nilSlice []int
	if nilSlice == nil {
		fmt.Printf("nilSlice is nil, its address is %p\n", nilSlice)
	}
	fmt.Println("append 10 to nilSlice")
	nilSlice = append(nilSlice, 10)
	if nilSlice == nil {
		fmt.Println("nilSlice is still nil")
	} else {
		fmt.Printf("nilSlice is not nil, %v its address is %p\n", nilSlice, nilSlice)
	}

	// make an empty slice, with capacity of 6
	slice2 := make([]int, 0, 6)
	if slice2 == nil {
		fmt.Println("slice2 is nil")
	} else {
		fmt.Printf("slice2 is not nil, its address is %p\n", slice2)

		for i := 1; i <= 4; i++ {
			fmt.Println("append", i*100, "to slice2:")
			slice2 = append(slice2, i*100)
			fmt.Printf("address of slice2 is %p, value of slice2 is %v\n", slice2, slice2)
		}
	}

	// slice2 has 4 elements now and its capacity is 6
	// append 1, 2, 3 to slice2, which will cause the slice to expand
	slice2 = append(slice2, []int{1, 2, 3}...)
	// slice2 has 7 elements now and its capacity is 12
	fmt.Printf("capacity of slice2: %v, length of slice2: %v, address of slice2: %p\n", cap(slice2), len(slice2), slice2)
	slice3 := slice2[2:6]
	// slice3 has 4 elements now and its capacity is 10 (caculation of capacity is 12 - 2 = 10)
	fmt.Println("capacity of slice3:", cap(slice3))
}
