/*
A={1,2,3}

Rank: 0 → Subset: []
Rank: 1 → Subset: [1]
Rank: 2 → Subset: [1, 2]
Rank: 3 → Subset: [2]
Rank: 4 → Subset: [2, 3]
Rank: 5 → Subset: [1, 2, 3]
Rank: 6 → Subset: [1, 3]
Rank: 7 → Subset: [3]
*/

package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// computes the rank of B
// in rank grey order.
func RankGrey(B, A []int) int {
    rank := 0
    for i, a := range A {
        if contains(B, a) {
            rank ^= (1 << i)
        }
    }
    return rank
}

// generates the subset of A corresponding 
// to the rank r in Grey order
func UnrankGrey(A []int, r int) []int {
    var subset []int
    for i := range A {
        if (r>>i)&1 == 1 {
            subset = append(subset, A[i])
        }
    }
    return subset
}

// Helper functions
func contains(set []int, elem int) bool {
    for _, v := range set {
        if v == elem {
            return true
        }
    }
    return false
}
func parseInput(input string) []int {
    var set []int
    for _, val := range strings.Split(input, " ") {
        var num int
        if _, err := fmt.Sscanf(val, "%d", &num); err == nil {
            set = append(set, num)
        }
    }
    return set
}

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Println("Enter the universal set A (space-separated integers):")
    inputA, _ := reader.ReadString('\n')
    A := parseInput(strings.TrimSpace(inputA))

    fmt.Println("Enter subset B (space-separated integers):")
    inputB, _ := reader.ReadString('\n')
    B := parseInput(strings.TrimSpace(inputB))

    rank := RankGrey(B, A)
    fmt.Printf("Rank of subset B: %d\n", rank)

    subset := UnrankGrey(A, rank)
    fmt.Printf("Subset with rank %d: %v\n", rank, subset)
}
