package main

import (
    "fmt"
)

func factorial(n int) int {
    // stop recursion when:
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
    var n int
    fmt.Print("Enter a positive integer: ")
    fmt.Scan(&n)

    // exit case
    if n < 1 {
        fmt.Println("Please enter a positive integer.")
        return
    }

    fmt.Printf("Factorial of %d is %d", n, factorial(n))
}
