/*
    - Prints all the permutations of given n.
*/
package main

import (
	"fmt"
)

func identityPerm(n int) []int {
	perm := make([]int, n)
	for i := 0; i < n; i++ {
		perm[i] = i + 1
	}
	return perm
}

func unrank(n int, r int, permIndex []int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		fact := factorial(n - 1 - i)
		index := r / fact
		result[i] = permIndex[index]
		// remove used element from pi
		permIndex = append(permIndex[:index], permIndex[index+1:]...)
		r %= fact
	}
	return result
}

func factorial(num int) int {
	if num == 0 || num == 1 {
		return 1
	}
	return num * factorial(num-1)
}

func main() {
    var n int

    fmt.Print("Enter the value of n: ")
    fmt.Scan(&n)

	for i := 0; i < factorial(n); i++ {
		p := identityPerm(n) // reset p for each r rank

        // i + 1 for 1- indexing the rank 
		fmt.Printf("Rank %d: %v\n", i + 1, unrank(n, i, p)) 
	}
}
