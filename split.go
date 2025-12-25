package main

// Split splits a string by the given separator and returns a slice of strings.
// It removes empty strings from the result (e.g., trailing separators).
func Split(s, sep string) []string {
	if sep == "" {
		return []string{s}
	}

	var result []string
	start := 0

	for i := 0; i <= len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			if start < i {
				result = append(result, s[start:i])
			}
			start = i + len(sep)
			i += len(sep) - 1
		}
	}

	// Add the remaining part
	if start < len(s) {
		result = append(result, s[start:])
	}

	// If the string ends with the separator, we've already handled it by not adding an empty string
	// But we need to make sure we don't add empty strings
	return result
}

