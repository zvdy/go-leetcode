func isAnagram(s string, t string) bool {
	// If the strings are not the same length, they cannot be anagrams.
	if len(s) != len(t) {
		return false
	}

	// Create a map to keep track of the characters in the first string.
	// The keys of the map are the characters, and the values are the number of times the character appears in the string.
	seenCharacters := make(map[rune]int)

	// Iterate through each character in the first string.
	for _, v := range s {
		// If the character already exists in the map, increment the count.
		if _, ok := seenCharacters[v]; ok {
			seenCharacters[v]++
		} else {
			// Otherwise, add the character to the map.
			seenCharacters[v] = 1
		}
	}

	// Iterate through each character in the second string.
	for _, v := range t {
		// If the character does not exist in the map, then it is not an anagram.
		if _, ok := seenCharacters[v]; !ok {
			return false
		}

		// If the character exists in the map, but the count is 0, then it is not an anagram.
		if seenCharacters[v] == 0 {
			return false
		}

		// Otherwise, decrement the count.
		seenCharacters[v]--
	}

	// If we have iterated through all the characters without finding a mismatch, then the strings are anagrams.
	return true
}