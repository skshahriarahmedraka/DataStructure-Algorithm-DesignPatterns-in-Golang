package main

import (
	"fmt"
	"math"
)

const V = 6 // Number of vertices

func floydWarshall(graph [][]int) {
	dist := make([][]int, V)

	// Copy the given graph into the distance matrix
	for i := range dist {
		dist[i] = make([]int, V)
		copy(dist[i], graph[i])
	}

	// Floyd-Warshall algorithm main loop
	for k := 0; k < V; k++ {
		for i := 0; i < V; i++ {
			for j := 0; j < V; j++ {
				// If vertex k is on the shortest path from i to j,
				// then update the value of dist[i][j]
				if dist[i][k] != math.MaxInt64 && dist[k][j] != math.MaxInt64 && dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	// Check for negative edge weight cycle
	for i := 0; i < V; i++ {
		if dist[i][i] < 0 {
			fmt.Println("Negative edge weight cycle is present")
			return
		}
	}

	// Print shortest path distances
	for i := 0; i < V; i++ {
		for j := 0; j < V; j++ {
			if dist[i][j] == math.MaxInt64 {
				fmt.Printf("INF ")
			} else {
				fmt.Printf("%d ", dist[i][j])
			}
		}
		fmt.Println()
	}
}

func main() {
	graph := [][]int{
		{0, 1, 4, math.MaxInt64, math.MaxInt64, math.MaxInt64},
		{math.MaxInt64, 0, 4, 2, 7, math.MaxInt64},
		{math.MaxInt64, math.MaxInt64, 0, 3, 4, math.MaxInt64},
		{math.MaxInt64, math.MaxInt64, math.MaxInt64, 0, math.MaxInt64, 4},
		{math.MaxInt64, math.MaxInt64, math.MaxInt64, 3, 0, math.MaxInt64},
		{math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, 5, 0},
	}

	floydWarshall(graph)
}
