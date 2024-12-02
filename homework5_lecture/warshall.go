package main

import (
	"fmt"
	"math"
)
/*
*
*  this program generates each matrix towards the final Shortest Path Matrix.
*  therefore you can see each step it takes to compute the final mtx.
*
*  presented to us in the course 10 of GTC. 
*/

// NOTES:
// The “shortest path” is defined as the path with the minimum sum of edge weights.
// Number of vertices in the graph
const vertices_number = 5
// lets write the matrix here of the given directed graph.
/*

0 10 50 65  ∞
∞  0 30  ∞  4
∞  ∞  0 20 44
∞ 70  ∞  0 23
6  ∞  ∞  ∞  0

*/
// we init the  weighted matrix here:
// instead of 0 for INF we are going to use math.MaxInt (then we can check it.)

var graph = [vertices_number][vertices_number]int{
	{0, 10, 50, 65, math.MaxInt},
	{math.MaxInt, 0, 30, math.MaxInt, 4},
	{math.MaxInt, math.MaxInt, 0, 20, 44},
	{math.MaxInt, 70, math.MaxInt, 0, 23},
	{6, math.MaxInt, math.MaxInt, math.MaxInt, 0},
}

func floydWarshall(graph [vertices_number][vertices_number]int) [vertices_number][vertices_number]int {
	// Create a distance matrix initialized to the graph values
	dist := graph

	// Floyd-Warshall Algorithm
    // computes AG* in O(n^3) 
    countMtx := 0;
	for k := 0; k < vertices_number; k++ {
		for i := 0; i < vertices_number; i++ {
			for j := 0; j < vertices_number; j++ {
				if dist[i][k] != math.MaxInt && dist[k][j] != math.MaxInt {
					dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
				}
			}
		}

        countMtx++;
        if(countMtx+1 > vertices_number){
            fmt.Printf("\n________________________\n");
            fmt.Printf("\nthis is the matrix no %d",countMtx);
	        fmt.Println("\nShortest Path Matrix:\n")
            printMatrix(dist);
        }else{
            fmt.Printf("\nthis is the matrix no %d\n\n",countMtx);
            printMatrix(dist);
        }
	}

	return dist
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func printMatrix(matrix [vertices_number][vertices_number]int) {
	for i := 0; i < vertices_number; i++ {
		for j := 0; j < vertices_number; j++ {
			if matrix[i][j] == math.MaxInt {
				fmt.Print("INF ")
			} else {
				fmt.Printf("%3d ", matrix[i][j])
			}
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("Original Graph:")
	printMatrix(graph)
    fmt.Printf("\n________________________\n");

    // computes and prints each matrix towards the final Shortest Path Matrix.
	floydWarshall(graph)
}
