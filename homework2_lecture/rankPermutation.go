/*3. int rankPermutation(int n, int* p) which takes as input arguments
a positive integer n and a (pointer to) an array of n integers, and computes the rank of the permutation hp[0], p[1], . . . , p[n − 1]i in lexicographic
order. We assume implicitly that hp[0], p[1], . . . , p[n − 1]i is a permutation
of 1,2,. . . ,n.*/
package main

import (
	"fmt"
)

func factorial(n int) int {
    result := 1
    for i := 2; i <= n; i++ { result *= i }
    return result
}
func isPermutation(n int, p []int) bool {
    seen := make([]bool, n+1)
    for _, num := range p {
        if num < 1 || num > n || seen[num] { return false }
        seen[num] = true
    }
    return true
}

// the rank is calculated by multiplying the count of smaller unseen numbers by
// the factorial of the remaining numbers. 
func rankPermutation(n int, p []int) int {
    rank := 0
    seen := make([]bool, n+1)

    // the seen unseen thingy. 
    for i := 0; i < n; i++ {
        count := 0
        for j := 1; j < p[i]; j++ {
            if !seen[j] {
                count++
            }
        }
        rank += count * factorial(n-i-1)
        seen[p[i]] = true
    }

    // 1based rank
    return rank + 1
}

func main() {
    var n int
    fmt.Print("Enter the value of n: ")
    fmt.Scan(&n)

    // exit case
    if n<1 {
        fmt.Print("n must equal grater than 0")
        return
    }

    p := make([]int, n)
    fmt.Printf("Enter %d integers (as a permutation of 1 to %d): ", n, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&p[i])
    }

    // exit case
    if !isPermutation(n, p) {
        fmt.Println("Input is not a valid permutation.")
        return
    }

    rank := rankPermutation(n, p)
    fmt.Printf("The rank of the permutation is: %d\n", rank)
}
