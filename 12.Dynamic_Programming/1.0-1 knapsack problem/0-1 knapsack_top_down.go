package main

import (
	"fmt"
	"github.com/k0kubun/pp/v3"
)

func main() {
    // Test case 1
    // MaxCapacity1 := 10
    // WeightList1 := []int{2, 3, 4, 5}
    // ValueList1 := []int{3, 4, 5, 6}
    // NumberOfValue1 := len(WeightList1)
    // fmt.Println("Test case 1:", KnapSack(MaxCapacity1, WeightList1, ValueList1, NumberOfValue1)) // Output: 7

    // // Test case 2
    // MaxCapacity2 := 50
    // WeightList2 := []int{10, 20, 30}
    // ValueList2 := []int{60, 100, 120}
    // NumberOfValue2 := len(WeightList2)
    // fmt.Println("Test case 2:", KnapSack(MaxCapacity2, WeightList2, ValueList2, NumberOfValue2)) // Output: 220

    // // Test case 3
    MaxCapacity3 := 5
    WeightList3 := []int{2, 1, 3}
    ValueList3 := []int{4, 2, 3}
    NumberOfValue3 := len(WeightList3) 
    fmt.Println("Test case 3:", KnapSack(MaxCapacity3, WeightList3, ValueList3, NumberOfValue3)) // Output: 7
}

func KnapSack(MaxCapacity int, WeightList []int, ValueList []int, NumberOfValue int) int {
	// Create a memoization table
	memo := make([][]int, NumberOfValue+1)
	for i := range memo {
		memo[i] = make([]int, MaxCapacity+1)
	}

	// Initialize the memoization table with initial values
	for i := 0; i <= NumberOfValue; i++ {
		for j := 0; j <= MaxCapacity; j++ {
			if i == 0 || j == 0 {
				memo[i][j] = 0
			}
		}
	}

	// Fill the memoization table iteratively
	for i := 1; i <= NumberOfValue; i++ {
		for j := 1; j <= MaxCapacity; j++ {
			if WeightList[i-1] <= j {
				fmt.Println(i, j, WeightList[i-1], j-WeightList[i-1], memo[i-1][j-WeightList[i-1]], memo[i-1][j])
				memo[i][j] = Max(ValueList[i-1]+memo[i-1][j-WeightList[i-1]], memo[i-1][j])
			} else {
				memo[i][j] = memo[i-1][j]
			}
		}
	}
	pp.Println(memo)
	// Return the maximum value from the memoization table
	return memo[NumberOfValue][MaxCapacity]
}

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
