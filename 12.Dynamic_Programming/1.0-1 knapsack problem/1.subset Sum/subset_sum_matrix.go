package main

import "fmt"

// Returns true if there is a subset of set[]
// with sum equal to given sum
func isSubsetSum(set []int, n, sum int) bool {
	// The value of subset[i][j] will be true if
	// there is a subset of set[0..j-1] with sum
	// equal to i
	subset := make([][]bool, n+1)
	for i := range subset {
		subset[i] = make([]bool, sum+1)
	}

	// If sum is 0, then answer is true
	for i := 0; i <= n; i++ {
		subset[i][0] = true
	}

	// If sum is not 0 and set is empty,
	// then answer is false
	for i := 1; i <= sum; i++ {
		subset[0][i] = false
	}

	// Fill the subset table in bottom up manner
	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			if j < set[i-1] {
				subset[i][j] = subset[i-1][j]
			}
			if j >= set[i-1] {
				subset[i][j] = subset[i-1][j] || subset[i-1][j-set[i-1]]
			}
		}
	}

	return subset[n][sum]
}

// Driver code
func main() {
	set := []int{3, 34, 4, 12, 5, 2}
	sum := 9
	n := len(set)
	if isSubsetSum(set, n, sum) {
		fmt.Println("Found a subset with given sum")
	} else {
		fmt.Println("No subset with given sum")
	}
}
