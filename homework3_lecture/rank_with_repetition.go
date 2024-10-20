package main

import (
	"fmt"
	"math"
)

// calculates the rank for given permutation with repetition.
func computeRank(permutation []int, n int) int {
	r := len(permutation)
	rank := 0
	for i := 0; i < r; i++ {
		// we calculate how many permutations can be formed 
        // with the remaining elements
		rank += permutation[i] * int(math.Pow(float64(n), float64(r-i-1)))
	}
	return rank
}

func main() {
	var n, r int

	fmt.Print("Enter the number of elements in the set (n): ")
	fmt.Scan(&n)

	fmt.Print("Enter the length of the permutation (r): ")
	fmt.Scan(&r)

	permutation := make([]int, r)
	fmt.Printf("Enter the permutation of length %d (space separated): ", r)
	for i := 0; i < r; i++ {
		fmt.Scan(&permutation[i])
	}

	rank := computeRank(permutation, n)
	fmt.Printf("The rank of permutation %v is: %d\n", permutation, rank)
}
