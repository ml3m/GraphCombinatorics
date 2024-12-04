package main

import (
	"fmt"
	"math"
	"container/heap"
)

const numNodes = 6 // Number of nodes in the graph

// Edge structure for adjacency list representation
type Edge struct {
	to, weight int
}

// PriorityQueueItem for implementing a min-heap
type PriorityQueueItem struct {
	node, distance int
}

// PriorityQueue is a min-heap of PriorityQueueItems
type PriorityQueue []*PriorityQueueItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*PriorityQueueItem)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// Dijkstra's Algorithm
func dijkstra(graph map[int][]Edge, source int) ([]int, []int) {
	// Distance array and predecessor array
	distance := make([]int, numNodes)
	predecessor := make([]int, numNodes)

	// Initialize distances to infinity and predecessors to -1
	for i := range distance {
		distance[i] = math.MaxInt
		predecessor[i] = -1
	}
	distance[source] = 0

	// Priority queue for selecting the next node with the smallest distance
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &PriorityQueueItem{node: source, distance: 0})

	// Process the graph
	for pq.Len() > 0 {
		current := heap.Pop(pq).(*PriorityQueueItem)
		currentNode := current.node

		// Explore neighbors
		for _, edge := range graph[currentNode] {
			newDist := distance[currentNode] + edge.weight
			if newDist < distance[edge.to] {
				distance[edge.to] = newDist
				predecessor[edge.to] = currentNode
				heap.Push(pq, &PriorityQueueItem{node: edge.to, distance: newDist})
			}
		}
	}

	return distance, predecessor
}

// Print the results of the algorithm
func printResults(distance []int, predecessor []int) {
	fmt.Println("Node\tDistance\tPredecessor")
	for i := 0; i < numNodes; i++ {
		pred := "None"
		if predecessor[i] != -1 {
			pred = fmt.Sprintf("%d", predecessor[i])
		}
		fmt.Printf("%d\t%d\t\t%s\n", i, distance[i], pred)
	}
}

func main() {
	// Graph representation as an adjacency list
	graph := map[int][]Edge{
        // s
        0: {{1,25}, {2,35}},

        // b
        1: {{3, 90}, {2, 15}},

        // c
        2: {{1, 50}, {3, 30}, {4, 50}},

        // e
        3: {{4, 10}, {5, 70}},

        // d
        4: {{3, 60}, {5, 20}},

        // f
        5: {},
	}

	source := 0 // Source node (s)

	// Compute shortest paths using Dijkstra's algorithm
	distance, predecessor := dijkstra(graph, source)

	// Print the results
	fmt.Println("Results of Dijkstra's Algorithm:")
	printResults(distance, predecessor) }
