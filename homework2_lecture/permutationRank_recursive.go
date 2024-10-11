package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func Rank(p []int) int {
	n := len(p)
	if n == 1 {
        // made it so exit condition returns 1 such that it goes from indexing
        // from 0 to indexing from 1 because that is the correct visual, rank
        // must be >= 1
		return 1
	}

	q := make([]int, n-1)

	// p[0 .. n-1] to become a permutation of {1, ..., n-1}
	for i := 1; i < n; i++ {
		q[i-1] = p[i]
		if q[i-1] > p[0] {
			q[i-1]--
		}
	}

	rank := Rank(q) + (p[0]-1)*factorial(n-1)
	return rank
}

func main() {
	var permutation []int
	var n int

	fmt.Print("Enter the number of elements in the permutation: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var num int
		fmt.Printf("Enter element %d: ", i+1)
		fmt.Scan(&num)
		permutation = append(permutation, num)
	}

	rank := Rank(permutation)
	fmt.Printf("The rank of the permutation %v is: %d\n", permutation, rank)
}
