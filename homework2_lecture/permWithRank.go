/*4. permWithRank(int k, int n) which prints the permutation with rank k
of {1, 2, . . . , n} in lexicographic order.package main*/
package main

import (
    "fmt"
)

func factorial(n int) int {
    result := 1
    for i := 2; i <= n; i++ { result *= i }
    return result
}

func permWithRank(k, n int) {
    rankDisplay := k;
    k-- // 0-based indexing
    numbers := make([]int, n)
    // numbers vector -> fill 1 - n 
    for i := 0; i < n; i++ {
        numbers[i] = i + 1 
    }

    result := []int{}
    used := make([]bool, n+1) // for track used numbers

    for i := 0; i < n; i++ {
        factorialValue := factorial(n - i - 1) // calculate factorial of remaining positions
        index := k / factorialValue //find the index of the current digit
        k %= factorialValue //updte k for the next position

        count := 0
        for j := 1; j <= n; j++ {
            if !used[j] {
                if count == index {
                    result = append(result, j)
                    used[j] = true // number used. 
                    break
                }
                count++
            }
        }
    }

    fmt.Printf("Permutation with rank %d is: ", rankDisplay)
    for _, num := range result {
        fmt.Printf("%d ", num)
    }
    fmt.Println()
}

func main() {
    var k, n int
    fmt.Print("Enter the value of n: ")
    fmt.Scan(&n)
    fmt.Print("Enter the rank k: ")
    fmt.Scan(&k)

    // exit case
    if n<=0 {
        fmt.Print("n must equal grater than 0")
        return
    }

    // exit case
    // Validate the rank k
    // make sure the rank that the user input is not higher than is possible
    // with chosen n
    if k < 1 || k > factorial(n) {
        fmt.Printf("Error: Rank k must be between 1 and %d (inclusive).\n", factorial(n))
        return
    }

    permWithRank(k, n)
}
