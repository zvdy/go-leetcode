func topKFrequent(nums []int, k int) []int {
	// create a map to count the frequency of each number
	numCounts := make(map[int]int)
	for _, num := range nums {
		numCounts[num]++
	}

	// create a slice of pairs (number, count) from the map
	numFreqPairs := make([][2]int, 0, len(numCounts))
	for num, count := range numCounts {
		numFreqPairs = append(numFreqPairs, [2]int{num, count})
	}

	// use quickselect to find the k most frequent numbers
	numFreqPairs = quickSelect(numFreqPairs, 0, len(numFreqPairs)-1, k-1)

	// create a slice of the k most frequent numbers
	result := make([]int, 0, k)
	for _, pair := range numFreqPairs[:k] {
		result = append(result, pair[0])
	}

	return result
}

func quickSelect(numFreqPairs [][2]int, left, right, k int) [][2]int {
	// choose the rightmost element as the pivot
	pivot := numFreqPairs[right]
	// partition the slice around the pivot
	j := left
	for i := left; i < right; i++ {
		if numFreqPairs[i][1] >= pivot[1] {
			numFreqPairs[i], numFreqPairs[j] = numFreqPairs[j], numFreqPairs[i]
			j++
		}
	}
	numFreqPairs[j], numFreqPairs[right] = numFreqPairs[right], numFreqPairs[j]

	// recursively partition the left or right side of the slice
	if j < k {
		return quickSelect(numFreqPairs, j+1, right, k)
	} else if j > k {
		return quickSelect(numFreqPairs, left, j-1, k)
	} else {
		return numFreqPairs
	}
}