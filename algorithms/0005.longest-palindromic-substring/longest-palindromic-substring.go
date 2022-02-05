package Problem0005

func longestPalindrome(s string) string {
	if len(s) < 2 { // must be a palindrome, return directly
		return s
	}

	// The index of the first character of the longest palindrome, and the length of the longest palindrome
	begin, maxLen := 0, 1

	// in for loop
	// b represents the **first** character index number of the palindrome,
	// e represents the **tail** character index number of the palindrome,
	// i represents the index number of the first character of the palindrome `middle section`
	// in each for loop
	// Start with i, use the same characteristics of all characters in the `middle section`, let b and e point to the beginning and end of the `middle section` respectively
	// Then expand from the `middle section` to both sides, let b and e point to the beginning and end of the longest palindrome centered on this `middle section`
	for i := 0; i < len(s); { // Start looking for the longest palindrome with s[i] as the first character of the `middle segment`.
		if len(s)-i <= maxLen/2 {
			// because i is the index number of the first character of the palindrome `middle section`
			// Assuming that the length of the longest palindrome that can be found at this time is l, then, l <= (len(s)-i)*2 - 1
			// if, len(s)-i <= maxLen/2 holds
			// then, l <= maxLen - 1
			// then, l < maxLen
			// So, no need to look any further.
			break
		}

		b, e := i, i
		for e < len(s)-1 && s[e+1] == s[e] {
			e++
			// After the loop ends, s[b:e+1] is the same string
		}

		// The first character of the `middle section` of the next palindrome will only be s[e+1]
		// prepare for the next loop
		i = e + 1

		for e < len(s)-1 && b > 0 && s[e+1] == s[b-1] {
			e++
			b--
			// After the loop ends, s[b:e+1] is the longest palindrome we can find this time.
		}

		newLen := e + 1 - b
		// If you create a new record, update the record
		if newLen > maxLen {
			begin = b
			maxLen = newLen
		}
	}

	return s[begin : begin+maxLen]
}