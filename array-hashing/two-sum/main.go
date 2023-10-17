func twoSum(nums []int, target int) []int {
	// Create a map to keep track of the numbers in the array.
	// The keys of the map are the numbers, and the values are the indices.
	seenNumbers := make(map[int]int)

	// Iterate through each number in the array.
	for i, v := range nums {
		// Calculate the complement.
		complement := target - v

		// If the complement exists in the map, then we have found the pair.
		if _, ok := seenNumbers[complement]; ok {
			// Return the indices of the pair.
			return []int{seenNumbers[complement], i}
		}

		// Otherwise, add the number to the map.
		seenNumbers[v] = i
	}

	// If we have iterated through all the numbers without finding a pair, then there is no solution.
	return nil
}