package main

import (
	"fmt"
)

// compute the next lexicographic permutation with repetition
func nextPermutation(permutation []int, n int) []int {
	r := len(permutation)
	next := make([]int, r)
	copy(next, permutation)

	for i := r - 1; i >= 0; i-- {
		if next[i] < n-1 {
			next[i]++
			return next
		} else {
			next[i] = 0
		}
	}

	return next // if it's the last permutation this will return the first one
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

	next := nextPermutation(permutation, n)
	fmt.Printf("The next permutation after %v is: %v\n", permutation, next)
}
