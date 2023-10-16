// containsDuplicate checks if a slice of integers contains any duplicates.
// It returns true if there are duplicates, false otherwise.
func containsDuplicate(nums []int) bool {
	// If the slice has less than 2 elements, there can be no duplicates.
	if len(nums) < 2 {
		return false
	}

	// Create a map to keep track of the integers in the slice.
	// The keys of the map are the integers, and the values are always true.
	seenIntegers := make(map[int]bool)

	// Iterate through each integer in the slice.
	for _, v := range nums {
		// If the integer already exists in the map, then there is a duplicate.
		if _, ok := seenIntegers[v]; ok {
			return true
		}
		// Otherwise, add the integer to the map.
		seenIntegers[v] = true
	}

	// If we have iterated through all the integers without finding a duplicate, then there are no duplicates.
	return false
}