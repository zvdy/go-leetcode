func productExceptSelf(nums []int) []int {
	// Initialize the result array
	result := make([]int, len(nums))
	// Initialize the left and right arrays
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	// Set the first element of the left array to 1
	left[0] = 1
	// Set the last element of the right array to 1
	right[len(nums)-1] = 1
	// Calculate the left array
	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	// Calculate the right array
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}
	// Calculate the result array
	for i := 0; i < len(nums); i++ {
		result[i] = left[i] * right[i]
	}
	return result
}