package main

import (
	"fmt"
	"math"
)

// we calculate the k-th r-permutation with repetition
func kthPermutation(k, n, r int) []int {
	permutation := make([]int, r)
	for i := 0; i < r; i++ {
		divisor := int(math.Pow(float64(n), float64(r-i-1)))
		permutation[i] = k / divisor
		k = k % divisor
	}
	return permutation
}

func main() {
	var n, r, k int

	fmt.Print("Enter the number of elements in the set (n): ")
	fmt.Scan(&n)
	fmt.Print("Enter the length of the permutation (r): ")
	fmt.Scan(&r)
	fmt.Print("Enter the rank (k): ")
	fmt.Scan(&k)

	result := kthPermutation(k, n, r)
	fmt.Printf("The permutation of rank %d is: %v\n", k, result)
}
