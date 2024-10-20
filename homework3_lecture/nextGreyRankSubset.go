package main

/*

Gray code: 000 → Subset: [] (Empty subset)
Gray code: 001 → Subset: [1]
Gray code: 011 → Subset: [1, 2]
Gray code: 010 → Subset: [2]
Gray code: 110 → Subset: [2, 3]
Gray code: 111 → Subset: [1, 2, 3]
Gray code: 101 → Subset: [1, 3]
Gray code: 100 → Subset: [3]

*/

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// computes the rank of subset B in Gray code order
func RankGrey(B, A []int) int {
    rank := 0
    for i, a := range A {
        if contains(B, a) {
            rank ^= (1 << i)
        }
    }
    return rank
}

// generates the subset of A 
// corresponding to the rank r in Gray code order
func UnrankGrey(A []int, r int) []int {
    var subset []int
    for i := range A {
        if (r>>i)&1 == 1 {
            subset = append(subset, A[i])
        }
    }
    return subset
}

// computes the next subset after B in Gray code order
func NextGreyRankSubset(A, B []int) []int {
    // get rank
    rank := RankGrey(B, A)

    // get next grey code
    nextRank := rank ^ (rank + 1)

    // return subset corresponding to the next Gray code rank
    return UnrankGrey(A, nextRank)
}


/* Helper functions */

// check if a subset contains an element
func contains(set []int, elem int) bool {
    for _, v := range set {
        if v == elem {
            return true
        }
    }
    return false
}

// parse input into a slice of integers
func parseInput(input string) []int {
    var set []int
    input = strings.TrimSpace(input)
    for _, val := range strings.Split(input, " ") {
        var num int
        fmt.Sscanf(val, "%d", &num)
        set = append(set, num)
    }
    return set
}

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Println("Enter the universal set A (space-separated integers):")
    inputA, _ := reader.ReadString('\n')
    A := parseInput(inputA)

    fmt.Println("Enter subset B (space-separated integers):")
    inputB, _ := reader.ReadString('\n')
    B := parseInput(inputB)

    nextSubset := NextGreyRankSubset(A, B)
    fmt.Printf("Next subset after B: %v\n", nextSubset)
}
