/* void isPermutation(int n, int* p) which takes as input arguments a positive
integer n and a (pointer to) an array of n integers, and prints the message •
"is a permutation" if hp[0], . . . , p[n − 1]i is a permutation of the numbers
1, . . . , n. • "not a permutation" otherwise. Write a program that reads from
console the values of p[0], . . . , p[n − 1] and then invokes
isPermutation(n,p) to detect if what was read is a permutation or not. package
main
*/
package main

import (
	"fmt"
)

func isPermutation(n int, p []int) {
    // track presence of numbers from 1 to n
    seen := make([]bool, n)

    for i := 0; i < n; i++ {
        // exit case
        // If the number is out of range or already seen therefore -> it's 
        // not a permutation
        if p[i] < 1 || p[i] > n || seen[p[i]-1] {
            fmt.Println("not a permutation")
            return
        }
        // mark as seen 
        seen[p[i]-1] = true
    }

    // this means all numbers from 1 - n are seen only once. 
    fmt.Println("is a permutation")
}

func main() {
    var n int
    fmt.Print("Enter the value of n: ")
    fmt.Scan(&n)

    // exit case
    if n<=0 {
        fmt.Print("n must equal grater than 0")
        return
    }

    // Create a slice of size n
    p := make([]int, n)
    fmt.Printf("Enter %d integers: ", n)
    for i := 0; i < n; i++ {
        fmt.Scan(&p[i])
    }

    isPermutation(n, p)
}
