package main

import "fmt"

// Returns true if there is a subset
// of set[] with sum equal to given sum
func isSubsetSum(set []int, n, sum int) bool {
	// Base Cases
	if sum == 0 {
		return true
	}
	if n == 0 {
		return false
	}

	// If last element is greater than sum,
	// then ignore it
	if set[n-1] > sum {
		return isSubsetSum(set, n-1, sum)
	}

	// Else, check if sum can be obtained by any
	// of the following:
	// (a) including the last element
	// (b) excluding the last element
	return isSubsetSum(set, n-1, sum) || isSubsetSum(set, n-1, sum-set[n-1])
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
