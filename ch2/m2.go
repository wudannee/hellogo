// Simple generics and mathstuffs with fibonacci, max
package ch2

import (
	"fmt"
	"math/big"

	"golang.org/x/exp/constraints"
)

type MathStuff[T constraints.Ordered] struct {
}

// Fibonacci calculates the nth Fibonacci number using big.Int to handle large numbers.
func (ms *MathStuff[T]) Fibonacci(n int) *big.Int {
	a := big.NewInt(0)
	b := big.NewInt(1)

	for i := 0; i < n; i++ {
		temp := new(big.Int).Set(b)
		b.Add(a, b)
		a.Set(temp)
	}

	return a
}

// Max returns the maximum of a and b.
func (ms *MathStuff[T]) Max(a, b T) T {
	if a > b {
		return a
	}
	return b
}

func DemoMathStuffs() {
	mathStuff := &MathStuff[float64]{}
	if mathStuff == nil {
		fmt.Println("mathStuff is nil")
	}

	result := mathStuff.Max(1.2, 3.14)
	fmt.Println("result:", result)
	for i := 0; i <= 3; i++ {
		fmt.Printf("fib(%d) = %s\n", i, mathStuff.Fibonacci(i).String())
	}
}
