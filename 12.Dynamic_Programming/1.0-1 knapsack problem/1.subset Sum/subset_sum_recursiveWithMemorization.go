package main

import "fmt"

// Taking the matrix as globally
var tab [2000][2000]int

// Check if possible subset with
// given sum is possible or not
func subsetSum(a []int, n, sum int) int {
	// If the sum is zero it means
	// we got our expected sum
	if sum == 0 {
		return 1
	}

	if n <= 0 {
		return 0
	}

	// If the value is not -1 it means it
	// already called the function
	// with the same value.
	// it will save our from the repetition.
	if tab[n-1][sum] != -1 {
		return tab[n-1][sum]
	}

	// If the value of a[n-1] is
	// greater than the sum.
	// we call for the next value
	if a[n-1] > sum {
		return tab[n-1][sum] = subsetSum(a, n-1, sum)
	} else {
		// Here we do two calls because we
		// don't know which value is
		// full-fill our criteria
		// that's why we doing two calls
		if subsetSum(a, n-1, sum) == 1 || subsetSum(a, n-1, sum-a[n-1]) == 1 {
			tab[n-1][sum] = 1
		} else {
			tab[n-1][sum] = 0
		}
		return tab[n-1][sum]
	}
}

// Driver Code
func main() {
	// Storing the value -1 to the matrix
	for i := 0; i < 2000; i++ {
		for j := 0; j < 2000; j++ {
			tab[i][j] = -1
		}
	}
	n := 5
	a := []int{1, 5, 3, 7, 4}
	sum := 12

	if subsetSum(a, n, sum) == 1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
