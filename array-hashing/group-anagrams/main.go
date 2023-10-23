// Key is a type alias for a fixed-length byte array of length 26 (the number of letters in the alphabet).
// Each byte in the array represents the count of a particular letter in a string.
type Key [26]byte

// strKey takes a string and returns a Key that represents the count of each letter in the string.
// It does this by iterating through each character in the string and incrementing the corresponding byte in the Key.
func strKey(str string) (key Key) {
	for i := range str {
		key[str[i]-'a']++
	}
	return
}

// groupAnagrams takes a slice of strings and groups them into anagrams.
// It does this by creating a map where the keys are Keys (representing the count of each letter in a string)
// and the values are slices of strings that have the same Key.
// It then iterates through the map and appends each slice of strings to a result slice.
func groupAnagrams(strs []string) [][]string {
	// Create a map to store the groups of anagrams.
	groups := make(map[Key][]string)

	// Iterate through each string in the input slice.
	for _, v := range strs {
		// Get the Key for the string.
		key := strKey(v)
		// Append the string to the slice of strings associated with the Key in the map.
		groups[key] = append(groups[key], v)
	}

	// Create a slice to store the result.
	result := make([][]string, 0, len(groups))
	// Iterate through each slice of strings in the map and append it to the result slice.
	for _, v := range groups {
		result = append(result, v)
	}
	return result
}