package main

import (
    "fmt"
)

/*
example:

a1 = 3 means that in the inverse perm., position 3 should have
the value 1 (b3 = 1)

a2 = 5 means that position 5 should have 
the value 2 (b4 = 3)

a3 = 4 means that position 4 should have 
the value 3 (b1 = 4)

a4 = 1 means that position 1 should have
the value 4

(...)

*/

func main() {
    var n int
    fmt.Print("Enter number of elements n: ")
    fmt.Scan(&n)

    a := make([]int, n)
    fmt.Printf("Enter %d elements of the permutation:\n", n)
    for i := 0; i < n; i++ {
        fmt.Scan(&a[i])
    }

    b := make([]int, n)

    for i := 0; i < n; i++ {
        b[a[i]-1] = i + 1
    }

    fmt.Println("The inverse permutation is:")
    for _, val := range b {
        fmt.Printf("%d ", val)
    }
    fmt.Println()
}
