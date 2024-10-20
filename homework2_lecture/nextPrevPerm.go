/*
5. void nextPermutation(int n, int* p) which takes as input arguments
a positive integer n and a (pointer to) an array of n integers, and displays the permutation next to hp[0], p[1], . . . , p[n − 1]i in lexicographic order. We assume implicitly that hp[0], p[1], . . . , p[n − 1]i is a permutation
of 1,2,. . . ,n.
*/
package main

import (
	"fmt"
	"sort"
)

<<<<<<< HEAD
// check if the slice is a valid permutation of 1 to n
func isValidPermutation(p []int, n int) bool {
	seen := make([]bool, n+1) // Using n+1 for 1-based index
	for _, num := range p {
		if num < 1 || num > n || seen[num] {
			return false // Not a valid permutation
		}
		seen[num] = true // Mark the number as seen
	}
	return true // It's a valid permutation
}

// find and display the next permutation
=======
// Function to find and display the next permutation
>>>>>>> 74d61dd (organisation)
func nextPermutation(n int, p []int) {
	// Step 1: Find the first element from the end that is smaller than its next element
	i := n - 2
	for i >= 0 && p[i] >= p[i+1] {
		i--
	}

	// If no such element exists, the permutation is the last one, so sort to get the smallest
	if i < 0 {
		sort.Ints(p)
	} else {
		// Step 2: Find the rightmost element that is greater than p[i]
		j := n - 1
		for p[j] <= p[i] {
			j--
		}

		// Step 3: Swap p[i] and p[j]
		p[i], p[j] = p[j], p[i]

		// Step 4: Sort the sequence from p[i + 1] to the end to get the next permutation
		sort.Ints(p[i+1:])
	}

	fmt.Print("Next permutation is: ")
	for _, v := range p {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

// find and display the previous permutation
func prevPermutation(n int, p []int) {
	// Step 1: Find the first element from the end that is greater than its next element
	i := n - 2
	for i >= 0 && p[i] <= p[i+1] {
		i--
	}

	// If no such element exists, the permutation is the first one, so sort to get the largest
	if i < 0 {
		sort.Sort(sort.Reverse(sort.IntSlice(p))) // Sort to get the largest permutation
	} else {
		// Step 2: Find the rightmost element that is smaller than p[i]
		j := n - 1
		for p[j] >= p[i] {
			j--
		}

		// Step 3: Swap p[i] and p[j]
		p[i], p[j] = p[j], p[i]

		// Step 4: Sort the sequence from p[i + 1] to the end to get the previous permutation
		sort.Sort(sort.Reverse(sort.IntSlice(p[i+1:])))
	}

	// Display the result
	fmt.Print("Previous permutation is: ")
	for _, v := range p {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Print("Enter the value of n: ")
	fmt.Scan(&n)
	if n < 1 {
		fmt.Print("n must be >= 1")
		return
	}

	p := make([]int, n)
	fmt.Printf("Enter %d integers (as a permutation of 1 to %d): ", n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&p[i])
	}

	if !isValidPermutation(p, n) {
		fmt.Println("Invalid input! Please enter a valid permutation of the numbers 1 to", n)
		return
	}

	nextPermutation(n, p)

	fmt.Printf("Enter the same %d integers again for previous permutation: ", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&p[i])
	}

	if !isValidPermutation(p, n) {
		fmt.Println("Invalid input! Please enter a valid permutation of the numbers 1 to", n)
		return
	}

	prevPermutation(n, p)
}
