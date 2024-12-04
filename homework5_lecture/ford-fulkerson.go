package main

import (
	"fmt"
	"math"
)

func dfs(capacity [][]int, flow [][]int, visited []bool, src, sink int, parent []int) bool {
	visited[src] = true
	if src == sink {
		return true
	}
	for v := 0; v < len(capacity); v++ {
		if !visited[v] && capacity[src][v]-flow[src][v] > 0 {
			parent[v] = src
			if dfs(capacity, flow, visited, v, sink, parent) {
				return true
			}
		}
	}
	return false
}

// Ford-Fulkerson Algorithm to find the maximum flow
func fordFulkerson(capacity [][]int, src, sink int) int {
	n := len(capacity)
	flow := make([][]int, n)
	for i := range flow {
		flow[i] = make([]int, n)
	}

	parent := make([]int, n)
	maxFlow := 0

	for {
		visited := make([]bool, n)
		if !dfs(capacity, flow, visited, src, sink, parent) {
			break
		}

		// Find the minimum residual capacity along the augmenting path
		augmentedFlow := math.MaxInt32
		for v := sink; v != src; v = parent[v] {
			u := parent[v]
			augmentedFlow = int(math.Min(float64(augmentedFlow), float64(capacity[u][v]-flow[u][v])))
		}

		// Update the flow and residual graph
		for v := sink; v != src; v = parent[v] {
			u := parent[v]
			flow[u][v] += augmentedFlow
			flow[v][u] -= augmentedFlow
		}

		maxFlow += augmentedFlow
	}

	return maxFlow
}

func main() {
	// Adjacency matrix representation of the graph
	// The graph corresponds to the one in the image
    // better representation needed.
	capacity := [][]int{
		{0, 13, 10, 0, 0, 0},
		{0, 0, 5, 24, 0, 0},
		{0, 0, 0, 15, 7, 0},
		{0, 0, 0, 0, 0, 9},
		{0, 0, 0, 0, 0, 16},
		{0, 0, 0, 0, 0, 0},
	}

	source := 0 // Index of the source node (s)
	sink := 5   // Index of the sink node (t)

	fmt.Printf("The maximum flow is %d\n", fordFulkerson(capacity, source, sink))
}
