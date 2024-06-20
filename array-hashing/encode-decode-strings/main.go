// Encodes a list of strings to a single string.
func encode(strs []string) (string, error) {
	if len(strs) < 0 || len(strs) >= 100 {
		return "", errors.New("Invalid number of strings")
	}
	for _, str := range strs {
		if len(str) < 0 || len(str) >= 200 {
			return "", errors.New("Invalid string length")
		}
	}
	return strings.Join(strs, "\x00"), nil
}

// Decodes a single string to a list of strings.
func decode(s string) ([]string, error) {
	strs := strings.Split(s, "\x00")
	if len(strs) < 0 || len(strs) >= 100 {
		return nil, errors.New("Invalid number of strings")
	}
	for _, str := range strs {
		if len(str) < 0 || len(str) >= 200 {
			return nil, errors.New("Invalid string length")
		}
	}
	return strs, nil
}