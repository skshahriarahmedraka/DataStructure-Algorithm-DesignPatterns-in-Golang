package main



func main() {


	
}



func KnapSack(MaxCapacity int, WeightList []int, ValueList []int, NumberOfValue int) int {
	// Create a memoization table
	memo := make([][]int, NumberOfValue+1)
	for i := range memo {
		memo[i] = make([]int, MaxCapacity+1)
	}

	// Initialize the memoization table with -1 to indicate that the subproblem hasn't been solved yet
	for i := 0; i <= NumberOfValue; i++ {
		for j := 0; j <= MaxCapacity; j++ {
			memo[i][j] = -1
		}
	}

	// Call the recursive function with memoization
	return knapSackMemo(MaxCapacity, WeightList, ValueList, NumberOfValue, memo)
}

func knapSackMemo(MaxCapacity int, WeightList []int, ValueList []int, NumberOfValue int, memo [][]int) int {
	// If the subproblem has already been solved, return the memoized value
	if memo[NumberOfValue][MaxCapacity] != -1 {
		return memo[NumberOfValue][MaxCapacity]
	}

	// Base case
	if MaxCapacity == 0 || NumberOfValue == 0 {
		return 0
	}

	// If the current item's weight exceeds the capacity, skip it
	if WeightList[NumberOfValue-1] > MaxCapacity {
		memo[NumberOfValue][MaxCapacity] = knapSackMemo(MaxCapacity, WeightList, ValueList, NumberOfValue-1, memo)
		return memo[NumberOfValue][MaxCapacity]
	}

	// Calculate the maximum value by either including or excluding the current item
	include := ValueList[NumberOfValue-1] + knapSackMemo(MaxCapacity-WeightList[NumberOfValue-1], WeightList, ValueList, NumberOfValue-1, memo)
	exclude := knapSackMemo(MaxCapacity, WeightList, ValueList, NumberOfValue-1, memo)

	// Update the memoization table with the maximum value
	memo[NumberOfValue][MaxCapacity] = Max(include, exclude)
	return memo[NumberOfValue][MaxCapacity]
}

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
